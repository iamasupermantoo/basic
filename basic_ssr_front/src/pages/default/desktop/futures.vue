<template>
    <div class="window-height">
        <div>future</div>
    </div>
</template>
  
<script lang="ts">
import { reactive, toRefs, defineComponent } from 'vue';

export default defineComponent({
    name: 'FutureIndex',
    setup() {
        const state = reactive({});

        return {
            ...toRefs(state),
        };
    },
});
</script>
  