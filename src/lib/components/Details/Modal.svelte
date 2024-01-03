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

<div class="modal-active fixed w-full h-full top-0 left-0 flex items-center justify-center z-50">
    <div class="modal-overlay absolute w-full h-full bg-white opacity-95"></div>
  
    <div class="modal-container fixed w-full h-full z-50 overflow-y-auto ">
      
  
    </div>
  </div>

<style>

</style>