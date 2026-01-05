<script lang="ts">
  interface Props {
    url: string;
    canGoBack: boolean;
    canGoForward: boolean;
    onGo: (url: string) => void;
    onBack: () => void;
    onForward: () => void;
  }

  let { url = $bindable(), canGoBack, canGoForward, onGo, onBack, onForward }: Props = $props();

  function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    onGo(url)
  }
</script>

<div class="address-bar">
  <div class="nav-buttons">
    <button
      onclick={onBack}
      disabled={!canGoBack}
      title="Back"
      aria-label="Go back"
    >
      Back
    </button>
    <button
      onclick={onForward}
      disabled={!canGoForward}
      title="Next"
      aria-label="Go forward"
    >
      Next
    </button>
  </div>
  <form onsubmit={handleSubmit}>
    <input
      type="text"
      bind:value={url}
      placeholder="gopher://gopher.floodgap.com"
      aria-label="URL"
    />
    <button type="submit">Go</button>
  </form>
</div>

<style>
  .address-bar {
    display: flex;
    gap: 0.5rem;
    align-items: center;
  }

  .nav-buttons {
    display: flex;
    gap: 0.25rem;
  }

  .nav-buttons button {
    background-color: #e0e0e0;
    color: #000000;
    border: 1px solid #999999;
    border-radius: 4px;
    padding: 0.4rem 1.2rem;
    cursor: pointer;
    font-size: 0.95rem;
    font-family: 'Agave Nerd Font Mono', monospace;
    transition: background-color 0.1s;
  }

  .nav-buttons button:hover:not(:disabled) {
    background-color: #d0d0d0;
  }

  .nav-buttons button:active:not(:disabled) {
    background-color: #c0c0c0;
  }

  .nav-buttons button:disabled {
    opacity: 0.3;
    cursor: not-allowed;
  }

  form {
    display: flex;
    gap: 0.5rem;
    flex: 1;
  }

  input {
    flex: 1;
    background-color: #ffffff;
    color: #000000;
    border: 1px solid #999999;
    border-radius: 4px;
    padding: 0.4rem 0.6rem;
    font-family: 'Agave Nerd Font Mono', monospace;
    font-size: 0.95rem;
  }

  input:focus {
    outline: 2px solid #0066cc;
    border-color: #0066cc;
  }

  button[type="submit"] {
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

  button[type="submit"]:hover {
    background-color: #d0d0d0;
  }

  button[type="submit"]:active {
    background-color: #c0c0c0;
  }
</style>