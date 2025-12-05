<script lang="ts">
  import { RSVP_Response, type Party } from '../../../../gen/guest/v1/guest_pb';
  import { formatNames } from '../format';
  import { client } from '../guest_client';
  import Squiggle from '../Squiggle.svelte';
  import Button from './Button.svelte';

  let {
    party = $bindable(),
    openDetails,
  }: {
    party: Party;
    openDetails: () => void;
  } = $props();

  async function updateRsvpResponse(response: RSVP_Response) {
    const updateResponse = await client.updatePartyRsvpResponse({
      name: party.name,
      response,
    });
    if (updateResponse.party) {
      party = updateResponse.party;
    }
  }

  function makeOptions(
    party: Party,
  ): Array<{ text: string; value: RSVP_Response }> {
    const acceptedText =
      party.guests.length > 1 ? 'Wir werden da sein' : 'Ich werde da sein';
    const declinedText =
      party.guests.length > 1
        ? 'Wir können leider nicht'
        : 'Ich kann leider nicht';
    return [
      { text: acceptedText, value: RSVP_Response.ACCEPTED },
      { text: declinedText, value: RSVP_Response.DECLINED },
    ];
  }
</script>

<div
  class="flex h-screen w-screen flex-col items-center gap-8 overflow-y-auto py-4 lg:justify-center"
>
  <div class="flex flex-col items-center justify-center gap-12">
    <div
      class="font-fontoon flex flex-col items-center gap-4 text-5xl text-orange-500"
    >
      <h1>Sarah &</h1>
      <h1>Sebastian</h1>
    </div>
    <div class="text-pink-500">
      <Squiggle>
        <div
          class="font-fontoon flex h-full flex-col items-center justify-center gap-4 p-4"
        >
          <h2 class="text-center text-2xl text-pink-500">
            Moin {formatNames(party)}!
          </h2>
          <div class="flex gap-4">
            {#each makeOptions(party) as option}
              <button
                class="cursor-pointer rounded-lg bg-pink-500 px-6 py-4 text-white transition-opacity hover:opacity-80"
                onclick={() => updateRsvpResponse(option.value)}
                >{option.text}</button
              >
            {/each}
          </div>
        </div>
      </Squiggle>
    </div>
    <div
      class="font-fontoon flex flex-col items-center gap-4 text-5xl text-orange-500"
    >
      <p>22.05.2026</p>
      <p>foodlab</p>
    </div>
  </div>
  <div class="lg:absolute lg:inset-x-0 lg:bottom-4">
    <div class="flex justify-center gap-2">
      <Button onclick={() => openDetails()}>Weitere Informationen</Button>
    </div>
  </div>
</div>
