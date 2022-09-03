<script>
    import { onMount } from "svelte";
import Result from "./Result.svelte";
    export let playlistID
    let genres = []
    let genreConditions = []
    let dateCondition = []
    let featureCondition = {}
    let startYear
    let Danceability
    let Instrumentalness
    let Valence
    const features = [{name:"Danceability", val:Danceability}, {name:"Instrumentalness", val:Instrumentalness}, {name:"Valence", val:Valence}]
    let loading = false
    let newPlaylist

    async function getGenreList() {
        await fetch("http://127.0.0.1:3000/api/genres").then((response) => {
            return response.json();
         }).then((data) => {
                genres = data
        });
    }

    async function sendData() {
        await fetch("http://127.0.0.1:3000/api/condition", {
            method: 'POST',
            body: JSON.stringify({playlist_ID: playlistID, Release_date:dateCondition, Genres: genreConditions, Features:featureCondition}),
            })
        loading = true
    }

    function handleSubmit(event) {
        collectDateCondition(event)
        collectFeatureDate(event)
        sendData()
    }

    function addGenre(event) {
        genreConditions.push(event.target.parentNode[0].value)
        genreConditions = genreConditions
    }

    function collectDateCondition(event) {
        dateCondition = []
        dateCondition.push(parseInt(event.target[2].value))
        dateCondition.push(parseInt(event.target[3].value))
        dateCondition = dateCondition
    }

    function collectFeatureDate(event) {
        let name = ""
        for (let index = 0; index < 3; index++) {
            let start = index*2+4
            let end = index*2+5
            if (event.target[start].value && event.target[end].value != "") {
                if (index == 0) {
                    name = "Danceability"
                }
                if (index == 1) {
                    name = "Instrumentalness"
                }
                if (index == 2) {
                    name = "Valence"
                }
                featureCondition[name] = [event.target[start].value/10, event.target[end].value/10]
            }
        }
    }

    onMount(getGenreList)
</script>




{#if loading==true}
<Result loading={loading}/>
{:else}
<form on:submit|preventDefault="{handleSubmit}">
    <label for="genre">Genre</label>
    <input type="text" list="genre-options" />
    <datalist id="genre-options">
        {#each genres as g}
            <option value={g} />
        {/each}
    </datalist>
    <button type="button" on:click="{addGenre}">add</button>

    {#each genreConditions as g}
        <li>{g}</li>
    {/each}

    <label for="date">Release Date Range</label>
    <input type="number" min="1000" max="2100" step="1" bind:value="{startYear}" placeholder="start"/>
    <input type="number" min={startYear} max="2100" step="1" placeholder="end"/>

    {#each features as f}
    <label for= {f.name}>{f.name}</label>
    <input type="number" min="0" max="10" step="1" bind:value="{f.val}" placeholder="start"/>
    <input type="number" min={f.val} max="10" step="1" placeholder="end"/>
    {/each}
    <button>start sorting</button>
</form>
{/if}