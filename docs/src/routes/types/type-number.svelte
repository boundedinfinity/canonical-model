<script lang="ts">
    import TypeCommon from './type-common.svelte'
    import { Min, Max, MultipleOf, InvalidList, ValidList } from "../properties";

    export let typeId: string
</script>

<TypeCommon {typeId} builtIn>
    <svelte:fragment slot="props">
        <li>
            <Min {typeId} />
        </li>
        <li>
            <Max {typeId} />
        </li>
        <li>
            <MultipleOf {typeId} />
        </li>
        <li>
            <InvalidList {typeId} />
        </li>
        <li>
            <ValidList {typeId} />
        </li>
    </svelte:fragment>
</TypeCommon>
