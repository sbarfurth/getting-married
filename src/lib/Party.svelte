<script lang="ts">
  import type { Party } from '../../gen/guest/v1/guest_pb';
  import ContactDataForm from './ContactDataForm.svelte';
  import { client } from './guest_client';

  const SCROLL_FULL = 200;
  const TOUCH_SCROLL_FACTOR = 0.7;

  let titleImage = import.meta.env.VITE_TITLE_IMAGE_URL;

  let { name }: { name: string } = $props();

  let scrollPos = $state(0);
  let overlayOpacity = $derived(scrollPos / SCROLL_FULL);
  let reverseOverlayOpacity = $derived(1 - overlayOpacity);
  let overlayTransform = $derived(`translateY(${30 - overlayOpacity * 30}px)`);

  async function getParty(name: string): Promise<Party> {
    const response = await client.getParty({ name });
    if (!response.party) {
      throw new Error('missing party in response');
    }

    return response.party;
  }

  function closeForm() {
    scrollFormTo(0);
  }

  function openForm() {
    scrollFormTo(SCROLL_FULL);
  }

  let scrollingForm = false;

  function scrollFormTo(target: number) {
    if (scrollingForm) return;
    scrollingForm = true;
    const openInterval = setInterval(() => {
      addScroll(target > scrollPos ? 2 : -2);
      if (scrollPos === target) {
        clearInterval(openInterval);
        scrollingForm = false;
      }
    }, 1);
  }

  function addScroll(delta: number) {
    scrollPos = Math.min(SCROLL_FULL, Math.max(0, scrollPos + delta));
  }

  document.addEventListener(
    'wheel',
    (event) => {
      addScroll(event.deltaY);
    },
    { passive: true },
  );

  let touchStartY: number | undefined;
  document.addEventListener(
    'touchstart',
    (event) => {
      touchStartY = event.touches[0].pageY;
    },
    { passive: true },
  );
  document.addEventListener(
    'touchmove',
    (event) => {
      if (touchStartY === undefined) return;
      const y = event.touches[0].pageY;
      const delta = touchStartY - y;
      addScroll(delta * TOUCH_SCROLL_FACTOR);
    },
    { passive: true },
  );
</script>

{#await getParty(name)}
  Gäste werden geladen...
{:then party}
  <div class="flex h-dvh w-dvw items-center justify-center overflow-hidden">
    <img
      src={titleImage}
      alt="Titelbild"
      class="pointer-events-none h-auto w-full md:h-full md:w-auto"
    />
    <div
      class={[
        'absolute inset-0 backdrop-blur-xs',
        overlayOpacity < 1 && 'pointer-events-none',
      ]}
      style:opacity={overlayOpacity}
    >
      <div class="bg-beige-500 absolute inset-0 z-10 opacity-80"></div>
      <div
        class="absolute inset-0 z-20 mx-2 flex items-center justify-center"
        style:transform={overlayTransform}
        style:opacity={overlayOpacity}
      >
        <ContactDataForm {party} close={closeForm}></ContactDataForm>
      </div>
    </div>
  </div>
  {#if !party.address.street && overlayOpacity < 1}
    <div
      class="absolute inset-x-4 bottom-4 z-20 flex justify-center text-pink-500"
      style:opacity={reverseOverlayOpacity}
    >
      <button
        class="bg-beige-500 flex cursor-pointer items-center rounded-lg border-2 border-orange-500 px-2 py-1 text-xs font-normal text-orange-500 shadow-md/20"
        onclick={openForm}
      >
        Kontaktdaten für die Einladung eingeben
        <span class="ml-2 animate-bounce">
          <i
            class="inline-block rotate-45 border-r border-b border-current p-0.5"
          ></i>
        </span>
      </button>
    </div>
  {/if}
{:catch err}
  Gäste konnten nicht geladen werden: {err.message}
{/await}
