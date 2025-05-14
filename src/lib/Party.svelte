<script lang="ts">
  import type { Party } from '../../gen/guest/v1/guest_pb';
  import ContactDataForm from './ContactDataForm.svelte';
  import { client } from './guest_client';

  const SCROLL_FULL = 200;
  const TOUCH_SCROLL_FACTOR = 0.4;

  let titleImage = import.meta.env.VITE_TITLE_IMAGE_URL;

  let { name }: { name: string } = $props();

  let scrollPos = $state(0);
  let overlayOpacity = $derived(scrollPos / SCROLL_FULL);
  let overlayTransform = $derived(`translateY(${30 - overlayOpacity * 30}px)`);

  async function getParty(name: string): Promise<Party> {
    const response = await client.getParty({ name });
    if (!response.party) {
      throw new Error('missing party in response');
    }

    return response.party;
  }

  function closeForm() {
    scrollPos = 0;
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
{:catch err}
  Gäste konnten nicht geladen werden: {err.message}
{/await}
