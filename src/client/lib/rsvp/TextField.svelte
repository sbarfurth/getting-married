<script lang="ts">
  import type { FormEventHandler } from 'svelte/elements';

  let {
    label,
    placeholder,
    value = $bindable(),
    onblur,
    oninput,
  }: {
    label?: string;
    placeholder?: string;
    value?: string;
    onblur?: (event: FocusEvent) => void;
    oninput?: FormEventHandler<HTMLTextAreaElement>;
  } = $props();
  const uid = $props.id();
</script>

<div class="flex flex-col items-center gap-2">
  {#if label}
    <label class="text-xl text-blue-500" for={uid}>{label}</label>
  {/if}

  <textarea
    id={uid}
    rows="1"
    class="w-80 resize-none rounded-lg border-2 border-blue-500 bg-blue-500/7 p-2 text-blue-500 placeholder:text-blue-500 placeholder:opacity-50 focus:border-blue-500 focus:outline-0"
    {placeholder}
    bind:value
    {onblur}
    {oninput}
  ></textarea>
</div>
