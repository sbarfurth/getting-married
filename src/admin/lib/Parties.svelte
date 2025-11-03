<script lang="ts">
  import {
    PartySchema,
    RSVP_Response,
    type Party,
  } from '../../../gen/guest/v1/guest_pb';
  import { client } from './guest_client';
  import { create } from '@bufbuild/protobuf';
  import { ConnectError } from '@connectrpc/connect';
  import PartyForm from './PartyForm.svelte';
  import SingleParty from './SingleParty.svelte';

  async function getParties(): Promise<readonly Party[]> {
    const response = await client.listParties({});
    return response.parties;
  }

  let parties = $state(getParties());
  let partyError = $state('');

  let editOrCreateParty = $state(
    create(PartySchema, {
      address: {},
      contact: {},
    }),
  );

  let selectedParty = $state<Party | undefined>(undefined);
  function togglePartySelected(party: Party): void {
    if (selectedParty?.name === party.name) {
      selectedParty = undefined;
    } else {
      selectedParty = party;
    }
  }

  async function createOrUpdateParty(
    event: SubmitEvent,
    party: Party,
  ): Promise<void> {
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
      closePartyDialog();
    } catch (e: unknown) {
      if (e instanceof ConnectError) {
        partyError = e.rawMessage;
      } else {
        partyError = String(e);
      }
    }
  }

  function createOrEditParty(partyToEdit?: Party): void {
    const dialog = document.getElementById(
      'party-form-dialog',
    ) as HTMLDialogElement;
    if (!dialog) {
      return;
    }
    if (partyToEdit) {
      editOrCreateParty = partyToEdit;
    }
    dialog.showModal();
  }

  function closePartyDialog(): void {
    const dialog = document.getElementById(
      'party-form-dialog',
    ) as HTMLDialogElement;
    if (!dialog) {
      return;
    }
    editOrCreateParty = create(PartySchema, {
      address: {},
      contact: {},
    });
    partyError = '';
    dialog.close();
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

  async function deleteRsvp(party: Party): Promise<void> {
    for (const guest of party.guests) {
      guest.rsvp = undefined;
    }
    await client.updateParty({
      party,
    });
    parties = getParties();
  }
</script>

<dialog
  id="party-form-dialog"
  class="bg-beige-500 m-auto w-2/3 rounded-lg border-2 border-orange-500 p-4"
>
  <header class="flex items-center justify-between text-orange-500">
    <h2 class="text-2xl">
      Partei {editOrCreateParty.name ? 'bearbeiten' : 'erstellen'}
    </h2>
    <button
      type="button"
      onclick={closePartyDialog}
      class="cursor-pointer rounded-lg p-2 text-xl leading-0 transition-colors hover:bg-orange-500/20"
    >
      <span class="material-symbols-outlined">close</span>
    </button>
  </header>
  <PartyForm
    party={editOrCreateParty}
    onSubmit={(event, party) => createOrUpdateParty(event, party)}
  />
</dialog>

<button
  class="absolute right-4 bottom-4 cursor-pointer rounded-full bg-orange-500 p-4 text-2xl leading-0 text-white shadow-xl transition-opacity hover:opacity-80"
  onclick={() => createOrEditParty()}
>
  <span class="material-symbols-outlined">add</span>
</button>

<main class="grid h-full grid-cols-[300px_1fr]">
  <aside class="border-r-2 border-orange-500">
    {#if partyError}
      Fehler: {partyError}
    {/if}

    {#await parties}
      Parteien werden geladen...
    {:then parties}
      <div class="flex flex-col gap-2 p-2">
        {#each parties as party}
          <button
            type="button"
            class={[
              'flex cursor-pointer items-center justify-between rounded-lg border border-blue-500 p-2 text-left transition-colors hover:bg-blue-500/20',
              selectedParty?.name === party.name ? 'bg-blue-500/30' : '',
            ]}
            onclick={() => togglePartySelected(party)}
          >
            <span>{party.displayName}</span>
            <span
              >({party.guests.filter(
                (guest) => guest.rsvp?.response === RSVP_Response.ACCEPTED,
              ).length}/{party.guests.length})</span
            >
          </button>
        {:else}
          <p>Keine Parteien vorhanden.</p>
        {/each}
      </div>
    {:catch err}
      Parteien konnten nicht geladen werden: {err.message}
    {/await}
  </aside>
  <div class="overflow-y-auto">
    {#if selectedParty}
      <SingleParty
        party={selectedParty}
        onEdit={(party) => createOrEditParty(party)}
        onDelete={(party) => deleteParty(party)}
        onDeleteRsvp={(party) => deleteRsvp(party)}
      />
    {:else}
      <div class="flex h-full items-center justify-center text-xl">
        Keine Partei ausgewählt.
      </div>
    {/if}
  </div>
</main>
