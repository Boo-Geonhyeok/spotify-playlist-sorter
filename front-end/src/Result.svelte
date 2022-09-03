<script>
import { onMount } from "svelte";
export let loading
let newPlaylist
let img
async function getData() {
    await fetch("http://127.0.0.1:3000/api/filter").then((response) => {
            return response.json();
            }).then((data) => {
                loading = false
                newPlaylist = data
                try {
                    img = newPlaylist["images"][0]["url"]
                } catch (error) {
                    img = "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBxISEhUSEhIVFRUVFRUVFRUVFRcVFRcVFRUXFxUVFRUYHSggGBolHRUVITEhJSkrLi4uFx8zODMtNygtLisBCgoKDQ0NDw0NDisZFRkrKzctLS0tKy0tLTctKy0tKy0tLS0rLTctNy0tNzcrLSstKy0tNystLS03Ny0rLS0rLf/AABEIAOEA4QMBIgACEQEDEQH/xAAWAAEBAQAAAAAAAAAAAAAAAAAAAQL/xAAbEAEBAQACAwAAAAAAAAAAAAAAARECQSGBwf/EABYBAQEBAAAAAAAAAAAAAAAAAAABAv/EABQRAQAAAAAAAAAAAAAAAAAAAAD/2gAMAwEAAhEDEQA/AMkBpCqgChIgALgIBQQDACgBixKsBFAAiKBDQANEBdNAEABNFwBQAAAUEAClAogBoigBADQUAqALqYuoCwCgiighQoIACiII0RNBVIQAMACwAASqCIoABoEAAWIAqABqomg0JS0FEKBANA0VABJAFVIApolACABiAKipAUomABgABAVABRAFKICiALAQAAFMDABbCwEAEAUVCKAhhgAABgQBIKAgpQRSQsBIoYCLYYAYACCggqRRSCkBBUAAoFIACKgAoAyqAohAUAFiKgBAgEUAQFBPAagKugCiKAgUDCiggAFMCArKpQEVAFRQBQEkWxMUCgAkUAChQTAxAbAAEKAKgKioCwADCwgBiYpQSwxUAkXABLFxKoJi4AFMAEWgAUQFQAUSKCCoAasNBFMANRSAAAFIYCKigCKARFgESmLQRQoAAIUoCiYAAAtQoABQUQBYIugAAVFAQCgAaBQAMVABaRAAUEoAINYAkoABQAAAABRACKSgJBUAFQAgAigAkXTRApqioFAAAVFwBDQBSosADUBcIIACgi1FABAUSgBQoFomqAIugCatAVAFEXQT0KAyolBRMAUADUVMBQANVFA0pYAhFARLVANQoABIAEUAEwFEAa0QEQFFEXUBQKBgGgIoAACoAAAJpQwEVAFwAAIAqCggEBQASLABDkALCfQEEAVb00AM0oCAgKLEAU5ACCAixQBIUAVOgBYkAUAEf//Z"
                }
                console.log(data);
            });
}
onMount(getData)
</script>

{#if loading == true}
    <h1>Loading...</h1>
{:else}
    <h1>New Sorted Playlist</h1>
    <div class="container">
        <img src={img} alt="" width="300" class="center" />
    </div>
    <div class="container">
        <a href="{newPlaylist.external_urls["spotify"]}">link</a>
    </div>
{/if}

<style>
    h1, img, a {
        text-align: center;
    }

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

        a {
            display: inline-block;
            text-align: center;
        }
</style>
