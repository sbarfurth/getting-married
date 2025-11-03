<script lang="ts">
  import { BUTTON } from '../../../styles/element_styles';
  import { RSVP_Response, type Party } from '../../../gen/guest/v1/guest_pb';
  import { formatNames } from '../../client/lib/format';

  let {
    party,
    onEdit = () => {},
    onDelete = () => {},
    onDeleteRsvp = () => {},
  }: {
    party: Party;
    onEdit?: (party: Party) => void;
    onDelete?: (party: Party) => void;
    onDeleteRsvp?: (party: Party) => void;
  } = $props();

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

  function hasRsvp(party: Party): boolean {
    return party.guests.some((guest) => !!guest.rsvp);
  }

  function formatResponse(response: RSVP_Response): string {
    switch (response) {
      case RSVP_Response.ACCEPTED:
        return 'Kommt';
      case RSVP_Response.DECLINED:
        return 'Kommt nicht';
      case RSVP_Response.UNSPECIFIED:
        return 'Unbekannt';
    }
  }
</script>

<div class="flex flex-col gap-2 p-2">
  <div class="flex items-center justify-between">
    <div>
      <h3 class="text-xl font-bold">{party.displayName}</h3>
      <p class="text-sm font-normal text-gray-500">{party.name}</p>
    </div>
    <div class="flex gap-2">
      <button
        onclick={() => onEdit(party)}
        class="cursor-pointer rounded-lg p-2 text-xl leading-0 text-orange-500 transition-colors hover:bg-orange-500/20"
        ><span class="material-symbols-outlined">edit</span></button
      >
      <button
        onclick={() => onDelete(party)}
        class="cursor-pointer rounded-lg p-2 text-xl leading-0 text-orange-500 transition-colors hover:bg-orange-500/20"
        ><span class="material-symbols-outlined">delete</span></button
      >
    </div>
  </div>
  <hr />
  <ul>
    {#each party.guests as guest}
      <li>
        <h4>{guest.firstName} {guest.lastName}</h4>
        {#if guest.rsvp}
          <div class="grid grid-cols-2 font-normal">
            <p>Antwort</p>
            <p>{formatResponse(guest.rsvp.response) || '-'}</p>
            <p>Essenspräferenzen</p>
            <p>{guest.rsvp.dietaryRestrictions || '-'}</p>
            <p>Musikwünsche</p>
            <p>{guest.rsvp.musicWishes || '-'}</p>
          </div>
        {/if}
      </li>
    {/each}
    {#if party.guests.length < party.maxGuests}
      <li>(+{party.maxGuests - party.guests.length})</li>
    {/if}
  </ul>
  <hr />
  <div class="grid grid-cols-[auto_1fr] gap-2">
    <label for="{party.name}-address">Adresse</label>
    <address class="font-normal not-italic" id="{party.name}-address">
      {#if party.address && party.address.street}
        {party.address.street}, {party.address.postalCode}
        {party.address.city}
      {:else}
        Keine Angabe.
      {/if}
    </address>

    <label for="{party.name}-email">E-Mail</label>
    <p class="font-normal" id="{party.name}-email">
      {#if party.contact && party.contact.email}
        {party.contact.email}
      {:else}
        Keine Angabe.
      {/if}
    </p>

    <label for="{party.name}-phone">Telefon</label>
    <p class="font-normal" id="{party.name}-phone">
      {#if party.contact && party.contact.phone}
        {party.contact.phone}
      {:else}
        Keine Angabe.
      {/if}
    </p>
  </div>
  <hr />
  <div class="grid grid-cols-[auto_1fr] gap-2">
    <label for="{party.name}-guest-management">Gästeverwaltung</label>
    <p class="font-normal" id="{party.name}-guest-management">
      {#if party.allowGuestSelfService}
        Ja
      {:else}
        Nein
      {/if}
    </p>
  </div>
  <hr />
  <div class="flex gap-2">
    <button onclick={() => copyLink(party)} class={BUTTON}>Link kopieren</button
    >
    <button onclick={() => copyMessage(party)} class={BUTTON}
      >Nachricht kopieren</button
    >
  </div>
  <div class="flex gap-2">
    {#if hasRsvp(party)}
      <button onclick={() => onDeleteRsvp(party)} class={BUTTON}
        >RSVP löschen</button
      >
    {/if}
  </div>
</div>
