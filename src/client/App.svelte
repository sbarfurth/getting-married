<script lang="ts">
  import { client } from './lib/guest_client';
  import Party from './lib/Party.svelte';
  import type { Party as PartyPb } from '../../gen/guest/v1/guest_pb';
  import Rsvp from './lib/rsvp/Rsvp.svelte';
  import { parseFeatureBoolean } from './lib/feature';

  const urlSearchParams = new URLSearchParams(window.location.search);
  const partyName = urlSearchParams.get('p') ?? '';

  const enableRsvp = parseFeatureBoolean(import.meta.env.VITE_ENABLE_RSVP);

  let party: PartyPb | undefined;
  let partyErr: string | undefined;

  async function initParty(name: string): Promise<void> {
    try {
      const response = await client.getParty({ name });
      if (!response.party) {
        throw new Error('missing party in response');
      }
      party = response.party;
    } catch (err: unknown) {
      if (err instanceof Error) {
        partyErr = err.message;
      } else {
        partyErr = String(err);
      }
    }
  }

  initParty(partyName);
</script>

{#if partyName}
  {#if party}
    {#if enableRsvp}
      <Rsvp bind:party />
    {:else}
      <Party {party} />
    {/if}
  {:else if partyErr}
    Gäste konnten nicht geladen werden: {partyErr}
  {:else}
    Gäste werden geladen...
  {/if}
{:else}
  <div class="flex h-svh w-screen flex-col items-center justify-center gap-4">
    <h1 class="text-main text-4xl font-bold">Hallo!</h1>
    <p class="font-normal">
      Bitte nutze den persönlichen Link, den du erhalten hast.
    </p>
  </div>
{/if}
