<script lang="ts">
  import Party from './lib/Party.svelte';

  const urlSearchParams = new URLSearchParams(window.location.search);
  const partyName = urlSearchParams.get('p') ?? '';
</script>

{#if partyName}
  <Party name={partyName} />
{:else}
  <div class="flex h-screen w-screen items-center justify-center">
    <h1 class="text-main text-4xl font-bold">Hallo!</h1>
  </div>
{/if}
