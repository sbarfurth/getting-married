<script lang="ts">
  import type { Guest } from '../../../../gen/guest/v1/guest_pb';
  import { INPUT, LABEL } from '../../../styles/element_styles';

  let { guest = $bindable() }: { guest: Guest } = $props();
</script>

<div class="mb-6 grid gap-6 md:grid-cols-2">
  <div>
    <label for="firstName" class={LABEL}>Vorname:</label>
    <input
      id="firstName"
      type="text"
      required
      bind:value={guest.firstName}
      placeholder="Max"
      class={INPUT}
    />
  </div>
  <div>
    <label for="lastName" class={LABEL}>Nachname:</label>
    <input
      id="lastName"
      type="text"
      required
      bind:value={guest.lastName}
      placeholder="Mustermann"
      class={INPUT}
    />
  </div>
</div>
