<script lang="ts">
  import {
    GuestSchema,
    PartySchema,
    type Party,
  } from '../../../gen/guest/v1/guest_pb';
  import { client } from '../lib/guest_client';
  import { LABEL, INPUT, BUTTON } from '../../../styles/element_styles';
  import { create } from '@bufbuild/protobuf';
  import GuestForm from '../lib/GuestForm.svelte';
  import { ConnectError } from '@connectrpc/connect';

  async function getParties(): Promise<readonly Party[]> {
    const response = await client.listParties({});
    return response.parties;
  }

  let parties = $state(getParties());

  let party = $state(
    create(PartySchema, {
      address: {},
      contact: {},
    }),
  );

  let partyError = $state('');

  function addGuest() {
    party.guests.push(create(GuestSchema));
  }

  function removeGuest() {
    party.guests.splice(-1, 1);
  }

  async function createOrUpdateParty(event: SubmitEvent): Promise<void> {
    event.preventDefault();
    try {
      if (party.name) {
        const response = await client.updateParty({
          party,
        });
        if (!response.party) {
          throw new Error('missing updated party in response');
        }
      } else {
        const response = await client.createParty({
          party,
        });
        if (!response.party) {
          throw new Error('missing created party in response');
        }
      }
      parties = getParties();
      resetParty();
    } catch (e: unknown) {
      if (e instanceof ConnectError) {
        partyError = e.rawMessage;
      } else {
        partyError = String(e);
      }
    }
  }

  function editParty(partyToEdit: Party): void {
    party = partyToEdit;
  }

  function resetParty(): void {
    party = create(PartySchema, {
      address: {},
      contact: {},
    });
    partyError = '';
  }

  async function deleteParty(party: Party): Promise<void> {
    if (!confirm(`Willst du die Partei ${party.name} wirklich löschen?`)) {
      return;
    }
    await client.deleteParty({
      name: party.name,
    });
    parties = getParties();
  }

  function generateLink(party: Party): string {
    const canonicalDomain = import.meta.env.VITE_CANONICAL_URL;
    const url = new URL(
      `${canonicalDomain || document.location.origin}?p=${party.name}`,
    );
    return url.toString();
  }

  function copyMessage(party: Party): Promise<void> {
    return copyText([
      `Hallo ${formatNames(party)}!`,
      '',
      'Wir heiraten und freuen uns, unser Save the Date zu teilen.',
      `Mehr Infos gibt es unter dem Link.`,
      '',
      'Liebe Grüße,',
      'Sarah und Sebastian',
    ]);
  }

  function copyLink(party: Party): Promise<void> {
    return copyText([generateLink(party)]);
  }

  async function copyText(lines: string[]): Promise<void> {
    await navigator.clipboard.writeText(lines.join('\r\n'));
    alert('Kopiert!');
  }

  function formatNames(party: Party): string {
    if (party.guests.length === 1) {
      return party.guests[0].firstName;
    }
    const commaGuests = party.guests.slice(0, -1);
    const andGuest = party.guests[party.guests.length - 1];
    return `${commaGuests.map((g) => g.firstName).join(', ')} und ${andGuest.firstName}`;
  }
</script>

<main>
  {#if partyError}
    Fehler: {partyError}
  {/if}
  <form class="m-4 max-w-3xl" onsubmit={createOrUpdateParty}>
    <div class="mb-6 grid gap-6 md:grid-cols-2">
      <div>
        <label for="displayName" class={LABEL}>Anzeigename:</label>
        <input
          id="displayName"
          type="text"
          required
          bind:value={party.displayName}
          placeholder="Max & Erika"
          class={INPUT}
        />
      </div>
      <div>
        <label for="maxGuests" class={LABEL}>Gästezahl:</label>
        <input
          id="maxGuests"
          type="number"
          min="0"
          required
          bind:value={party.maxGuests}
          placeholder="Gästezahl"
          class={INPUT}
        />
      </div>
      {#if party.address}
        <div>
          <label for="street" class={LABEL}>Straße:</label>
          <input
            id="street"
            type="text"
            bind:value={party.address.street}
            placeholder="Musterstraße"
            class={INPUT}
          />
        </div>
        <div>
          <label for="city" class={LABEL}>Stadt:</label>
          <input
            id="city"
            type="text"
            bind:value={party.address.city}
            placeholder="Musterstadt"
            class={INPUT}
          />
        </div>
        <div>
          <label for="postalCode" class={LABEL}>PLZ:</label>
          <input
            id="postalCode"
            type="text"
            bind:value={party.address.postalCode}
            placeholder="12345"
            class={INPUT}
          />
        </div>
      {/if}
      {#if party.contact}
        <div>
          <label for="email" class={LABEL}>E-Mail:</label>
          <input
            id="email"
            type="email"
            bind:value={party.contact.email}
            placeholder="max@mustermann.de"
            class={INPUT}
          />
        </div>
        <div>
          <label for="phone" class={LABEL}>Telefon:</label>
          <input
            id="phone"
            type="tel"
            bind:value={party.contact.phone}
            placeholder="+49123456789"
            class={INPUT}
          />
        </div>
      {/if}
    </div>
    <div>
      <button
        type="button"
        class={BUTTON}
        onclick={addGuest}
        disabled={party.guests.length >= party.maxGuests}
        >Gast hinzufügen</button
      >
      <button
        type="button"
        class={BUTTON}
        onclick={removeGuest}
        disabled={party.guests.length === 0}>Gast entfernen</button
      >
      {#each party.guests as guest}
        <GuestForm {guest}></GuestForm>
      {/each}
    </div>

    <button type="submit" class={BUTTON}>
      {#if party.name}
        Partei speichern
      {:else}
        Partei erstellen
      {/if}
    </button>

    {#if party.name}
      <button type="button" class={BUTTON} onclick={resetParty}>
        Abbrechen
      </button>
    {/if}
  </form>

  {#await parties}
    Parteien werden geladen...
  {:then parties}
    <div class="m-4 grid gap-3 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
      {#each parties as party}
        <div class="flex flex-col gap-2 rounded border-2 border-pink-500 p-2">
          <div>
            <h3 class="text-xl font-bold">{party.displayName}</h3>
            <p class="text-sm font-normal text-gray-500">{party.name}</p>
          </div>
          <hr />
          <ul>
            {#each party.guests as guest}
              <li>{guest.firstName} {guest.lastName}</li>
            {/each}
            {#if party.guests.length < party.maxGuests}
              <li>(+{party.maxGuests - party.guests.length})</li>
            {/if}
          </ul>
          <hr />
          <div>
            <h4>Adresse</h4>
            <p class="font-normal">
              {#if party.address && party.address.street}
                {party.address.street}, {party.address.postalCode}
                {party.address.city}
              {:else}
                Keine Angabe.
              {/if}
            </p>
          </div>
          <hr />
          <div>
            <h4>E-Mail</h4>
            <p class="font-normal">
              {#if party.contact && party.contact.email}
                {party.contact.email}
              {:else}
                Keine Angabe.
              {/if}
            </p>
          </div>
          <hr />
          <div>
            <h4>Telefon</h4>
            <p class="font-normal">
              {#if party.contact && party.contact.phone}
                {party.contact.phone}
              {:else}
                Keine Angabe.
              {/if}
            </p>
          </div>
          <div class="flex gap-2">
            <button onclick={() => copyLink(party)} class={BUTTON}
              >Link kopieren</button
            >
            <button onclick={() => copyMessage(party)} class={BUTTON}
              >Nachricht kopieren</button
            >
          </div>
          <div class="flex gap-2">
            <button onclick={() => editParty(party)} class={BUTTON}
              >Bearbeiten</button
            >
            <button onclick={() => deleteParty(party)} class={BUTTON}
              >Löschen</button
            >
          </div>
        </div>
      {:else}
        <p>Keine Parteien vorhanden.</p>
      {/each}
    </div>
  {:catch err}
    Parteien konnten nicht geladen werden: {err.message}
  {/await}
</main>
