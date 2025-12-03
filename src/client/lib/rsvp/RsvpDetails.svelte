<script lang="ts">
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
    update: () => void;
  } = $props();
</script>

<TextField
  label="Essenspräferenzen"
  placeholder="z.B. Allergien, vegetarisch/vegan, etc."
  bind:value={dietaryRestrictions}
  onblur={() => update()}
/>
<TextField
  label="Musikwünsche"
  bind:value={musicWishes}
  onblur={() => update()}
/>
<TextField label="Sonstiges" bind:value={note} onblur={() => update()} />
