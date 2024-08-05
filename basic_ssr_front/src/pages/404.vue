<template>
  <q-layout class="row items-center justify-center">
    <q-page-container>
      <div class="row justify-center q-mb-lg">
        <q-img no-spinner src="/images/404.png" width="200px" height="200px" />
      </div>
      <div class="text-center text-weight-bold q-mb-md text-h4">
        Page Not Found!!
      </div>
      <div class="row justify-center q-mt-xl">
        <div class="text-center text-color-6 col-md-10 col-xs-10">
          we’re sorry,the page you requested could not be found Please go back
          to the home page
        </div>
      </div>
    </q-page-container>
  </q-layout>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
export default defineComponent({
  name: 'ErrorIndex',
  setup() {
    return {};
  },
});
</script>
<style lang="scss" scoped></style>
