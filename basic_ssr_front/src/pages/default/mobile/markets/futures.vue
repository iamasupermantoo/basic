<template>
    <div class="window-height">
        <div>MarketsFutures</div>
    </div>
</template>

<script lang="ts">
import { reactive, toRefs, defineComponent } from 'vue';

export default defineComponent({
    name: 'MarketsFutures',
    setup() {
        const state = reactive({});

        return {
            ...toRefs(state),
        };
    },
});
</script>
