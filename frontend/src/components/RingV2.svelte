<script lang="ts">
  export let diameter;
  export let duration;
  export let color;
  export let strokeWidth = 20;

  let circumference = 2 * Math.PI * (diameter / 2);
  let progress = 0;
</script>

<div
  id="ring"
  style={`--max-dash-length: ${circumference}; --duration: ${duration}s; --stroke-width:${strokeWidth}; --color:${color}; --progress: ${progress}; --diameter: ${diameter + strokeWidth}px`}
>
  <svg
    width={diameter + strokeWidth}
    height={diameter + strokeWidth}
    viewBox="0 0 {diameter + strokeWidth} {diameter + strokeWidth}"
  >
    <circle
      class="background-ring"
      cx={(diameter + strokeWidth) / 2}
      cy={(diameter + strokeWidth) / 2}
      r={diameter / 2}
    />
    <circle
      class="progress-ring"
      cx={(diameter + strokeWidth) / 2}
      cy={(diameter + strokeWidth) / 2}
      r={diameter / 2}
      marker-end="url(#end-marker)"
    />
  </svg>
  <div class="leading-circle"></div>
</div>

<style>
  #ring {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    background: rgba(255, 192, 203, 0.247);
  }
  svg {
    transform: rotate(-90deg);
  }

  circle {
    fill: none;
    stroke-width: var(--stroke-width);
    /* stroke-linecap: round; */
  }

  .background-ring {
    stroke: rgba(255, 255, 255, 0.2);
  }

  .progress-ring {
    stroke: var(--color);
    stroke-dasharray: var(--max-dash-length);
    stroke-dashoffset: var(--max-dash-length);
    animation-name: spin;
    animation-duration: var(--duration);
    animation-iteration-count: infinite;
    animation-timing-function: linear;
  }

  .leading-circle {
    position: absolute;
    height: var(--diameter);
    width: 20px;
    top: 0px;
    left: 50%;
    margin-left: -10px;
    animation: rotateClockwise 60s linear infinite;
  }

  .leading-circle:before {
    background: linear-gradient(90deg, transparent 0 40%, var(--color) 40% 100%);
    /* background: var(--color); */
    display: block;
    content: '';
    height: 20px;
    width: 20px;
    border-radius: 50%;
    box-shadow: 3px 2px 3px rgba(0, 0, 0, 0.1);
  }

  @keyframes spin {
    0% {
      stroke-dashoffset: var(--max-dash-length);
    }
    100% {
      stroke-dashoffset: 0;
    }
  }

  @keyframes rotateClockwise {
    100% {
      transform: rotate(360deg);
    }
  }
</style>
