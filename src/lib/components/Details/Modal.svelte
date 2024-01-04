<script lang="ts">
    import { goto, preloadData, pushState } from '$app/navigation';
    import { page } from '$app/stores';

    export let plant: any = {};

    const fetchPlantDetails = async (pid: string) => {
        const plant = await import(`../../../lib/data/plants/${pid}.json`)
        console.log(plant.animals)
    };

    $: plant && fetchPlantDetails(plant)

</script>

