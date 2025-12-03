<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    children,
    color = 'blue',
    size = 'default',
    type = 'button',
    onclick,
  }: {
    children: Snippet;
    color?: 'blue' | 'pink';
    size?: 'small' | 'default';
    type?: 'submit' | 'button';
    onclick?: (event: MouseEvent) => void;
  } = $props();
</script>

<button
  {onclick}
  {type}
  class={[
    color === 'blue' && 'border-blue-500 text-blue-500',
    color === 'pink' && 'border-pink-500 text-pink-500',
    size === 'small' && 'px-4 py-1',
    size === 'default' && 'px-4 py-2 text-xl',
    'cursor-pointer rounded-lg border-2 leading-none transition-opacity hover:opacity-80',
  ]}>{@render children()}</button
>
