<script lang="ts">
  interface Props {
    prompt: string;
    onSubmit: (query: string) => void;
    onCancel: () => void;
  }

  let { prompt, onSubmit, onCancel }: Props = $props();
  let query = $state("");
  let inputElement: HTMLInputElement;

  function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    if (query.trim()) {
      onSubmit(query.trim());
    }
  }

  function handleKeydown(e: KeyboardEvent) {
    if (e.key === "Escape") {
      onCancel();
    }
  }

  $effect(() => {
    inputElement?.focus();
  });
</script>

<div class="overlay" onkeydown={handleKeydown} role="dialog" aria-modal="true">
  <div class="dialog">
    <div class="prompt">{prompt}</div>
    <form onsubmit={handleSubmit}>
      <input
        type="text"
        bind:this={inputElement}
        bind:value={query}
        placeholder="Enter search query..."
        aria-label="Search query"
      />
      <div class="buttons">
        <button type="button" onclick={onCancel}>Cancel</button>
        <button type="submit">Search</button>
      </div>
    </form>
  </div>
</div>

<style>
  .overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    align-items: center;
    justify-content: center;
    z-index: 1000;
  }

  .dialog {
    background-color: #ffffff;
    border: 1px solid #999999;
    border-radius: 8px;
    padding: 1.5rem;
    min-width: 400px;
    max-width: 90vw;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  }

  .prompt {
    margin-bottom: 1rem;
    font-size: 1rem;
    color: #333333;
  }

  form {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }

  input {
    background-color: #ffffff;
    color: #000000;
    border: 1px solid #999999;
    border-radius: 4px;
    padding: 0.5rem 0.75rem;
    font-family: 'Agave Nerd Font Mono', monospace;
    font-size: 0.95rem;
    width: 100%;
    box-sizing: border-box;
  }

  input:focus {
    outline: 2px solid #0066cc;
    border-color: #0066cc;
  }

  .buttons {
    display: flex;
    justify-content: flex-end;
    gap: 0.5rem;
  }

  button {
    background-color: #e0e0e0;
    color: #000000;
    border: 1px solid #999999;
    border-radius: 4px;
    padding: 0.4rem 1.2rem;
    cursor: pointer;
    font-family: 'Agave Nerd Font Mono', monospace;
    font-size: 0.95rem;
    transition: background-color 0.1s;
  }

  button:hover {
    background-color: #d0d0d0;
  }

  button:active {
    background-color: #c0c0c0;
  }

  button[type="submit"] {
    background-color: #0066cc;
    color: #ffffff;
    border-color: #0055aa;
  }

  button[type="submit"]:hover {
    background-color: #0055aa;
  }

  button[type="submit"]:active {
    background-color: #004499;
  }
</style>