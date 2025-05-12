<script lang="ts">
  import type { Party } from '../../../../gen/guest/v1/guest_pb';
  import ContactDataForm from './ContactDataForm.svelte';
  import { client } from './guest_client';

  const SCROLL_FULL = 200;

  let { name }: { name: string } = $props();

  let wheelPos = $state(0);
  let overlayOpacity = $derived(wheelPos / SCROLL_FULL);

  async function getParty(name: string): Promise<Party> {
    const response = await client.getParty({ name });
    if (!response.party) {
      throw new Error('missing party in response');
    }

    return response.party;
  }

  function closeForm() {
    wheelPos = 0;
  }

  document.addEventListener('wheel', (event) => {
    wheelPos = Math.min(SCROLL_FULL, Math.max(0, wheelPos + event.deltaY));
  });
</script>

{#await getParty(name)}
  Gäste werden geladen...
{:then party}
  <div class="flex h-screen w-screen items-center justify-center">
    <img
      src="./assets/title.png"
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
      <div class="absolute inset-0 z-20 mx-2 flex items-center justify-center">
        <ContactDataForm {party} close={closeForm}></ContactDataForm>
      </div>
    </div>
  </div>
{:catch err}
  Gäste konnten nicht geladen werden: {err.message}
{/await}
