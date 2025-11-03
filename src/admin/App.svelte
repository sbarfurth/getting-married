<script lang="ts">
  import { BUTTON, INPUT } from '../../styles/element_styles';
  import { clearToken, getToken, setToken } from './lib/auth';
  import Parties from './lib/Parties.svelte';

  let password = $state('');
  let loggedIn = $state(getToken() !== undefined);

  async function login(event: SubmitEvent) {
    event.preventDefault();
    try {
      const rawResponse = await fetch(import.meta.env.VITE_API_URL + 'login', {
        method: 'post',
        body: JSON.stringify({ password: password }),
      });
      if (rawResponse.ok) {
        const response = await rawResponse.json();
        setToken(response['access_token']);
        loggedIn = true;
      } else {
        throw new Error(await rawResponse.text());
      }
    } catch (e: unknown) {
      alert(`Login fehlgeschlagen: ${(e as Error).message}`);
    }
  }

  function logout() {
    clearToken();
    loggedIn = false;
  }
</script>

<div class="flex h-screen flex-col">
  <header
    class="flex items-center justify-between bg-orange-500 px-4 py-2 text-white"
  >
    <h1 class="text-xl">Gästeverwaltung</h1>
    {#if loggedIn}
      <button
        type="button"
        onclick={logout}
        class="cursor-pointer rounded-lg p-2 text-xl leading-0 transition-colors hover:bg-white/20"
      >
        <span class="material-symbols-outlined">logout</span>
      </button>
    {/if}
  </header>

  <div class="min-h-0 flex-1">
    {#if loggedIn}
      <Parties />
    {:else}
      <div class="flex justify-center pt-8">
        <form class="flex w-md flex-col gap-2" onsubmit={login}>
          <label for="password">Admin-Passwort</label>
          <input
            class={INPUT}
            type="password"
            name="password"
            id="password"
            bind:value={password}
          />
          <button class={BUTTON} type="submit">Anmelden</button>
        </form>
      </div>
    {/if}
  </div>
</div>
