<script lang="ts">
  import { type Party, RSVP_Response } from '../../../../gen/guest/v1/guest_pb';
  import { client } from '../guest_client';
  import Button from './Button.svelte';
  import RsvpNo from './RsvpNo.svelte';
  import RsvpSelection from './RsvpSelection.svelte';
  import RsvpYes from './RsvpYes.svelte';

  let { party = $bindable() }: { party: Party } = $props();

  function getPartyRsvp(party: Party): RSVP_Response {
    for (const guest of party.guests) {
      if (guest.rsvp?.response === RSVP_Response.ACCEPTED) {
        return RSVP_Response.ACCEPTED;
      }
      if (!guest.rsvp || guest.rsvp.response === RSVP_Response.UNSPECIFIED) {
        return RSVP_Response.UNSPECIFIED;
      }
    }
    return RSVP_Response.DECLINED;
  }

  let rsvp = $derived(getPartyRsvp(party));

  function openDetails() {
    const dialog = document.getElementById(
      'wedding-detail-dialog',
    ) as HTMLDialogElement;
    dialog.showModal();
  }

  function closeDetails() {
    const dialog = document.getElementById(
      'wedding-detail-dialog',
    ) as HTMLDialogElement;
    dialog.close();
  }
</script>

{#if rsvp === RSVP_Response.UNSPECIFIED}
  <RsvpSelection bind:party />
{:else if rsvp === RSVP_Response.DECLINED}
  <RsvpNo bind:party />
{:else}
  <RsvpYes bind:party />
{/if}

<dialog
  id="wedding-detail-dialog"
  class="bg-beige-500 m-auto w-2/3 rounded-lg border-2 border-blue-500 p-4"
>
  <header class="flex items-center justify-between">
    <h2 class="font-fontoon text-2xl text-blue-500">Details zur Hochzeit</h2>
    <Button onclick={() => closeDetails()}>Schließen</Button>
  </header>
</dialog>

<div class="absolute right-4 bottom-4">
  <Button onclick={() => openDetails()}>Details zur Hochzeit</Button>
</div>
