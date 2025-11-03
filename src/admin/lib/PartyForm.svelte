<script lang="ts">
  import { create } from '@bufbuild/protobuf';
  import { GuestSchema, type Party } from '../../../gen/guest/v1/guest_pb';
  import { BUTTON, INPUT, LABEL } from '../../../styles/element_styles';
  import GuestForm from './GuestForm.svelte';

  let {
    party = $bindable(),
    onSubmit = () => {},
  }: {
    party: Party;
    onSubmit?: (event: SubmitEvent, party: Party) => void;
  } = $props();

  function addGuest() {
    party.guests.push(create(GuestSchema));
  }

  function removeGuest() {
    party.guests.splice(-1, 1);
  }
</script>

<form onsubmit={(event) => onSubmit(event, party)}>
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
  <div class="flex items-center gap-2">
    <label
      for="plus-one-self-service"
      class="block text-sm font-medium text-gray-900">Gästeverwaltung:</label
    >
    <input
      id="plus-one-self-service"
      type="checkbox"
      class="cursor-pointer"
      bind:checked={party.allowGuestSelfService}
    />
  </div>
  <div class="mt-2 mb-2">
    <button
      type="button"
      class={BUTTON}
      onclick={addGuest}
      disabled={party.guests.length >= party.maxGuests}>Gast hinzufügen</button
    >
    <button
      type="button"
      class={BUTTON}
      onclick={removeGuest}
      disabled={party.guests.length === 0}>Gast entfernen</button
    >
    {#each party.guests as guest}
      <GuestForm {guest} />
    {/each}
  </div>

  <button type="submit" class={BUTTON}>
    {#if party.name}
      Partei speichern
    {:else}
      Partei erstellen
    {/if}
  </button>
</form>
