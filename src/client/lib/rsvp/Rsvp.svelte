<script lang="ts">
  import { type Party, RSVP_Response } from '../../../../gen/guest/v1/guest_pb';
  import Details from './Details.svelte';
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
  <RsvpSelection bind:party {openDetails} />
{:else if rsvp === RSVP_Response.DECLINED}
  <RsvpNo bind:party {openDetails} />
{:else}
  <RsvpYes bind:party {openDetails} />
{/if}

<dialog
  id="wedding-detail-dialog"
  class="bg-beige-500 m-auto h-full w-full overflow-hidden rounded-lg border-2 border-blue-500 lg:min-h-6/12 lg:w-2/3"
>
  <div class="flex h-full flex-col">
    <Details
      onclose={() => closeDetails()}
      showInnerCircle={party.innerCircle}
    />
  </div>
</dialog>
