<script lang="ts">
  import { create } from '@bufbuild/protobuf';
  import {
    GuestSchema,
    RSVP_Response,
    RSVPSchema,
    type Party,
    type RSVP,
  } from '../../../../gen/guest/v1/guest_pb';
  import { client } from '../guest_client';
  import Button from './Button.svelte';
  import TextField from './TextField.svelte';

  let { party = $bindable() }: { party: Party } = $props();

  async function switchResponseToNo() {
    const updateResponse = await client.updatePartyRsvpResponse({
      name: party.name,
      response: RSVP_Response.DECLINED,
    });
    if (updateResponse.party) {
      party = updateResponse.party;
    }
  }

  async function updateGuestRsvp(
    guestIndex: number,
    updater?: (rsvp: RSVP) => void,
  ) {
    const rsvp = party.guests[guestIndex].rsvp ?? create(RSVPSchema);
    if (updater) {
      updater(rsvp);
    }
    party.guests[guestIndex].rsvp = rsvp;
    const updateResponse = await client.updateGuestRsvp({
      name: party.name,
      index: guestIndex,
      rsvp: rsvp,
    });
    if (updateResponse.party) {
      party = updateResponse.party;
    }
  }

  async function updateGuestRsvpResponse(
    guestIndex: number,
    response: RSVP_Response,
  ) {
    updateGuestRsvp(guestIndex, (rsvp) => {
      rsvp.response = response;
    });
  }

  let addingGuest = $state(false);
  let firstName = $state('');
  let lastName = $state('');

  function doneAddingGuest() {
    addingGuest = false;
    firstName = '';
    lastName = '';
  }

  async function addGuest() {
    const updateResponse = await client.addGuest({
      name: party.name,
      guest: create(GuestSchema, {
        firstName,
        lastName,
      }),
    });
    if (updateResponse.party) {
      party = updateResponse.party;
    }
    doneAddingGuest();
  }

  async function removeGuest(guestIndex: number) {
    const updateResponse = await client.removeGuest({
      name: party.name,
      index: guestIndex,
    });
    if (updateResponse.party) {
      party = updateResponse.party;
    }
  }
</script>

<div class="font-fontoon flex h-screen w-screen items-center justify-center">
  <div class="flex flex-col items-center gap-20">
    <h1 class="text-5xl text-pink-500">
      Wir freuen uns auf {party.guests.length > 1 ? 'euch' : 'dich'}!
    </h1>
    <div class="flex gap-16">
      {#each party.guests as guest, index}
        <div class="flex flex-col items-center gap-8">
          <h3 class="text-3xl text-orange-500">
            {guest.firstName}
            {guest.lastName}
          </h3>
          <div class="flex w-full">
            <button
              onclick={() =>
                updateGuestRsvpResponse(index, RSVP_Response.ACCEPTED)}
              class={[
                'flex-1 cursor-pointer rounded-l-lg border-2 border-r border-blue-500 p-2 text-blue-500 transition-colors hover:bg-blue-500/10',
                guest.rsvp?.response === RSVP_Response.ACCEPTED &&
                  'bg-blue-500 text-white hover:bg-blue-500/90',
              ]}>Kommt</button
            >
            <button
              onclick={() =>
                updateGuestRsvpResponse(index, RSVP_Response.DECLINED)}
              class={[
                'flex-1 cursor-pointer rounded-r-lg border-2 border-l border-blue-500 p-2 text-blue-500 transition-colors hover:bg-blue-500/10',
                guest.rsvp?.response === RSVP_Response.DECLINED &&
                  'bg-blue-500 text-white hover:bg-blue-500/90',
              ]}>Kommt nicht</button
            >
          </div>
          {#if guest.rsvp}
            <TextField
              label="Essenspräferenzen"
              placeholder="z.B. Allergien, vegetarisch/vegan, etc."
              bind:value={guest.rsvp.dietaryRestrictions}
              onblur={() => updateGuestRsvp(index)}
            />
            <TextField
              label="Musikwünsche"
              bind:value={guest.rsvp.musicWishes}
              onblur={() => updateGuestRsvp(index)}
            />
            <TextField
              label="Sonstiges"
              bind:value={guest.rsvp.note}
              onblur={() => updateGuestRsvp(index)}
            />
          {/if}
          {#if party.allowGuestSelfService && index !== 0}
            <Button color="pink" size="small" onclick={() => removeGuest(index)}
              >+1 entfernen</Button
            >
          {/if}
        </div>
      {/each}
    </div>
    {#if party.allowGuestSelfService && party.guests.length < party.maxGuests}
      <form
        class="flex flex-col items-center gap-2 text-pink-500"
        onsubmit={() => addGuest()}
      >
        Du kannst noch eine weitere Person hinzufügen.
        {#if addingGuest}
          <div class="flex gap-2">
            <input
              class="flex-1 rounded-lg border-2 border-pink-500 p-2 text-pink-500 hover:bg-pink-500/5 focus:border-pink-500 focus:outline-0"
              type="text"
              required
              bind:value={firstName}
              placeholder="Vorname"
            />
            <input
              class="flex-1 rounded-lg border-2 border-pink-500 p-2 text-pink-500 hover:bg-pink-500/5 focus:border-pink-500 focus:outline-0"
              type="text"
              required
              bind:value={lastName}
              placeholder="Nachname"
            />
          </div>
          <div class="flex w-full justify-between">
            <Button color="pink" size="small" onclick={() => doneAddingGuest()}
              >abbrechen</Button
            >
            <Button color="pink" size="small" type="submit">hinzufügen</Button>
          </div>
        {:else}
          <Button color="pink" size="small" onclick={() => (addingGuest = true)}
            >+1 hinzufügen</Button
          >
        {/if}
      </form>
    {/if}
  </div>
</div>

<div class="absolute bottom-4 left-4">
  <Button onclick={() => switchResponseToNo()} color="pink"
    >{party.guests.length > 1
      ? 'Wir können doch nicht kommen'
      : 'Ich kann doch nicht kommen'}</Button
  >
</div>
