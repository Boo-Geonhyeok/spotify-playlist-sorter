<script>
import { onMount } from "svelte";
import Sort from "./Sort.svelte";

    let authURL

    async function getAuthURL() {
        await fetch("http://127.0.0.1:3000/api/url").then((response) => {
            return response.json();
         }).then((data) => {
                authURL = data.url;
        });
    }

    onMount(getAuthURL)
</script>

{#if sessionStorage.getItem("isRedirected") == "true"}
    <Sort />
{:else}
    <h1>You should be authorized by Spotify first. Click link.</h1>
    <div>
        <a href={authURL} on:click="{() => {sessionStorage.setItem("isRedirected", true)}}">Authorization</a>
    </div>
{/if}

<style>
    h1, div {
        text-align: center;
    }
</style>