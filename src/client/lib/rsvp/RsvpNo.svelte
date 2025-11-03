<script lang="ts">
  import { RSVP_Response, type Party } from '../../../../gen/guest/v1/guest_pb';
  import { formatNames } from '../format';
  import { client } from '../guest_client';
  import Button from './Button.svelte';

  let { party = $bindable() }: { party: Party } = $props();

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

<div class="font-fontoon flex h-screen w-screen items-center justify-center">
  <h1 class="max-w-1/2 text-center text-5xl leading-relaxed text-pink-500">
    Schade, dass {party.guests.length > 1
      ? 'ihr nicht kommt'
      : 'du nicht kommst'}
    {formatNames(party)} :(
  </h1>
</div>

<div class="absolute bottom-4 left-4">
  <Button onclick={() => switchResponseToYes()} color="pink"
    >{party.guests.length > 1 ? 'Wir können' : 'Ich kann'} doch kommen</Button
  >
</div>
