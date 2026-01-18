<script lang="ts">
  import type { RSVP } from 'gen/guest/v1/guest_pb';
  import TextField from './TextField.svelte';

  let {
    dietaryRestrictions = $bindable(),
    musicWishes = $bindable(),
    note = $bindable(),
    update,
  }: {
    dietaryRestrictions: string;
    musicWishes: string;
    note: string;
    update: (updater: (rsvp: RSVP) => void) => void;
  } = $props();

  const DEBOUNCE_MS = 500;

  function debounce(fn: () => void, delay: number) {
    let timeout: ReturnType<typeof setTimeout>;
    return () => {
      clearTimeout(timeout);
      timeout = setTimeout(fn, delay);
    };
  }
</script>

<TextField
  label="Essenspräferenzen"
  placeholder="z.B. Allergien, vegetarisch/vegan, etc."
  bind:value={dietaryRestrictions}
  oninput={debounce(
    () =>
      update((rsvp) => {
        rsvp.dietaryRestrictions = dietaryRestrictions;
      }),
    DEBOUNCE_MS,
  )}
/>
<TextField
  label="Musikwünsche"
  bind:value={musicWishes}
  oninput={debounce(
    () =>
      update((rsvp) => {
        rsvp.musicWishes = musicWishes;
      }),
    DEBOUNCE_MS,
  )}
/>
<TextField
  label="Sonstiges"
  bind:value={note}
  oninput={debounce(
    () =>
      update((rsvp) => {
        rsvp.note = note;
      }),
    DEBOUNCE_MS,
  )}
/>
