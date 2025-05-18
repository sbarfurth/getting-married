<script lang="ts">
  import type { Party } from '../../gen/guest/v1/guest_pb';
  import { BUTTON, INPUT } from '../../styles/element_styles';
  import { client } from './guest_client';

  let confirmationImage = import.meta.env.VITE_CONFIRMATION_IMAGE_URL;

  let { party, close }: { party: Party; close: () => void } = $props();

  let street = $state(party.address?.street ?? '');
  let postalCode = $state(party.address?.postalCode ?? '');
  let city = $state(party.address?.city ?? '');
  let email = $state(party.contact?.email ?? '');
  let phone = $state(party.contact?.phone ?? '');

  let sending = $state(false);
  let closing = $state(false);

  async function save(event: SubmitEvent) {
    event.preventDefault();
    sending = true;
    try {
      await client.updatePartyContactInfo({
        name: party.name,
        address: {
          street,
          postalCode,
          city,
        },
        contact: {
          email,
          phone,
        },
      });
      closing = true;
      setTimeout(() => {
        close();
        closing = false;
      }, 8000);
    } catch (e: unknown) {
      alert('Speichern fehlgeschlagen :(');
    } finally {
      sending = false;
    }
  }
</script>

{#if closing}
  <div class="flex flex-col items-center gap-2">
    {#if confirmationImage}
      <img
        src={confirmationImage}
        alt="Bild zur Bestätigung"
        class="h-auto w-10/12 max-w-80"
      />
    {/if}
    <h1 class="mt-2 text-xl text-orange-500">Abgeschickt</h1>
    <p class="font-medium text-blue-500">Bis bald!</p>
  </div>
{:else}
  <form class="m-4 flex w-sm max-w-full flex-col gap-4" onsubmit={save}>
    <h1 class="mb-2 text-center text-2xl text-orange-500">
      Kontaktdaten für die Einladung
    </h1>
    <div class="mb-4 text-center font-semibold text-blue-500">
      {#each party.guests as guest}
        <p>{guest.firstName} {guest.lastName}</p>
      {/each}
      {#if party.guests.length < party.maxGuests}
        <p class="font-normal">+{party.maxGuests - party.guests.length}</p>
      {/if}
    </div>
    <div>
      <label for="street" class="text-blue-500">Straße</label>
      <input
        class={INPUT}
        type="text"
        id="street"
        name="street"
        disabled={sending}
        bind:value={street}
      />
    </div>
    <div class="flex gap-4">
      <div class="flex-1/3">
        <label for="postal" class="text-blue-500">PLZ</label>
        <input
          class={INPUT}
          type="text"
          id="postal"
          name="postal"
          disabled={sending}
          bind:value={postalCode}
        />
      </div>
      <div class="flex-2/3">
        <label for="city" class="text-blue-500">Ort</label>
        <input
          class={INPUT}
          type="text"
          id="city"
          name="city"
          disabled={sending}
          bind:value={city}
        />
      </div>
    </div>
    <div class="mt-8">
      <label for="email" class="text-blue-500">E-Mail</label>
      <input
        class={INPUT}
        type="email"
        id="email"
        name="email"
        disabled={sending}
        bind:value={email}
      />
    </div>
    <div>
      <label for="phone" class="text-blue-500">Telefon</label>
      <input
        class={INPUT}
        type="tel"
        id="phone"
        name="phone"
        disabled={sending}
        bind:value={phone}
      />
    </div>
    <button class={BUTTON} type="submit"
      >{sending ? 'Moment...' : 'Abschicken'}</button
    >
  </form>
{/if}
