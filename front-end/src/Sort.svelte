<script>
    import { onMount } from "svelte";
import { append } from "svelte/internal";
    import SortForm from "./SortForm.svelte";
        let playlistID = ""
        let playlists = []
        let showCondition = false
        let info = "Choose playlist to sort"
        async function getPlaylist() {
            await fetch("http://127.0.0.1:3000/api/playlist").then((response) => {
                return response.json();
             }).then((data) => {
                console.log(data[0]);
                for (let index = 0; index < data.length; index++) {
                    let playlist = {
                    ...data[index],
                    size: 300
                    }
                    playlists.push(playlist)
                }
                playlists=playlists
            });
        }

        function handleClick(p) {
            playlists.forEach(element => {
                element.size = 100
            });
            p.size = 300
            playlists = playlists
            playlistID = p.id
            showCondition = true
            info = "Submit conditions. You don't have to fill it all."
        }
        onMount(getPlaylist)
    </script>

    <h1>{info}</h1>
    <div class="container">
    {#each playlists as p}
        <img src={p["images"][0]["url"]} alt="" width={p.size} class="center" on:click="{handleClick(p)}">
    {/each}
    </div>

    {#if showCondition == true}
        <SortForm playlistID={playlistID} />
    {/if}

    <style>
        .center {
            text-align:center;
            display: inline-block;
            margin-left: 5px;
            margin-right: 5px;
        }

        .container {
            display: flex;
            justify-content: center;
            align-items: center;
        }

        h1 {
            text-align: center;
        }
    </style>