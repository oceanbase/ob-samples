<script>
    import '../app.css';
    import { user } from '../stores/authStore.js';
    import Auth from '../components/Auth.svelte';
    import Navbar from '../components/Navbar.svelte';
    import { loadTodos } from '../stores/todoStore.js';
</script>

<div class="container mx-auto my-6 max-w-lg">
    {#if $user}
        <Navbar />
        <slot />
    {:else}
        <Auth />
    {/if}
</div>
