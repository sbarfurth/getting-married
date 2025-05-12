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
      console.log('failed to log in:', e);
    }
  }

  function logout() {
    clearToken();
    loggedIn = false;
  }
</script>

{#if loggedIn}
  <button class={BUTTON} type="button" onclick={logout}>Abmelden</button>
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
