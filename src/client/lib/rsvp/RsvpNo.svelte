<script lang="ts">
  import { RSVP_Response, type Party } from '../../../../gen/guest/v1/guest_pb';
  import { formatNames } from '../format';
  import { client } from '../guest_client';
  import Button from './Button.svelte';

  let {
    party = $bindable(),
    openDetails,
  }: {
    party: Party;
    openDetails: () => void;
  } = $props();

  async function switchResponseToYes() {
    const updateResponse = await client.updatePartyRsvpResponse({
      name: party.name,
      response: RSVP_Response.ACCEPTED,
    });
    if (updateResponse.party) {
      party = updateResponse.party;
    }
  }
</script>

<div
  class="font-fontoon flex h-screen w-screen flex-col items-center justify-center gap-4 p-2"
>
  <h1
    class="text-center text-3xl leading-relaxed text-pink-500 lg:max-w-1/3 lg:text-5xl"
  >
    Schade, dass {party.guests.length > 1
      ? 'ihr nicht kommt'
      : 'du nicht kommst'}
    {formatNames(party)} :(
  </h1>
  <div class="flex gap-2">
    <Button onclick={() => openDetails()} size="small"
      >Weitere Informationen</Button
    >
  </div>
  <button
    class="cursor-pointer text-sm text-pink-500 underline"
    onclick={() => switchResponseToYes()}
    >{party.guests.length > 1 ? 'Wir können' : 'Ich kann'} doch kommen</button
  >
</div>
