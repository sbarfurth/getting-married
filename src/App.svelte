<script lang="ts">
  import Party from './lib/Party.svelte';

  const urlSearchParams = new URLSearchParams(window.location.search);
  const partyName = urlSearchParams.get('p') ?? '';
</script>

{#if partyName}
  <Party name={partyName} />
{:else}
  <div
    class="flex h-screen w-screen flex-col items-center justify-center gap-4"
  >
    <h1 class="text-main text-4xl font-bold">Hallo!</h1>
    <p class="font-normal">
      Bitte nutze den persönlichen Link, den du erhalten hast.
    </p>
  </div>
{/if}
