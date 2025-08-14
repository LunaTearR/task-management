<template>
  <div>
    <apexchart
      type="line"
      :options="chartOptions"
      :series="chartSeries"
    />
  </div>
</template>

<script lang="ts" setup>
import ApexCharts from 'vue3-apexcharts'

defineProps({
  chartOptions: {
    type: Object,
    required: true,
  },
  chartSeries: {
    type: Array,
    required: true,
  },
})

const chartOptions = {
  chart: {
    id: 'vuechart-example',
    height: 350,
    toolbar: {
      show: false,
    },
  },
  xaxis: {
    categories: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul'],
  },
  title: {
    text: 'Task Completion Progress',
    align: 'center',
    style: {
      fontSize: '18px',
      fontWeight: 'bold',
    },
  },
}

const chartSeries = [
  {
    name: 'Completed Tasks',
    data: [30, 40, 35, 50, 49, 60, 70],
  },
]
</script>

<script setup>
import { defineComponent } from "vue";

defineComponent({
  components: {
    apexchart: ApexCharts
  },
});
</script>

