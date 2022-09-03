<script>
    import { onMount } from "svelte";
    import SortForm from "./SortForm.svelte";
    let playlistID = ""
        let playlists = []
        let showCondition = false
        let info = "choose playlist to sort"
        async function getAuthURL() {
            await fetch("http://127.0.0.1:3000/api/playlist").then((response) => {
                return response.json();
             }).then((data) => {
                    playlists = data
            });
        }

        function handleClick(p) {
            playlists = [p]
            playlistID = p.id
            showCondition = true
            info = "Submit conditions. You don't have to fill it all."
        }
        onMount(getAuthURL)
    </script>

    <h1>{info}</h1>
    {#each playlists as p}
        <img src={p["images"][0]["url"]} alt="" width="300" on:click="{handleClick(p)}">
    {/each}

    {#if showCondition == true}
        <SortForm playlistID={playlistID}/>
    {/if}