import { useEffect, useMemo, useState } from 'react';
import Particles, { initParticlesEngine } from '@tsparticles/react';
// import { OutMode } from '@tsparticles/engine';
// import { loadAll } from "@tsparticles/all"; // if you are going to use `loadAll`, install the "@tsparticles/all" package too.
// import { loadFull } from "tsparticles"; // if you are going to use `loadFull`, install the "tsparticles" package too.
import { loadSlim } from '@tsparticles/slim'; // if you are going to use `loadSlim`, install the "@tsparticles/slim" package too.
// import { loadBasic } from "@tsparticles/basic"; // if you are going to use `loadBasic`, install the "@tsparticles/basic" package too.

import './particles.css';

const ParticlesDemo = ({ color = '#0d47a1', init = false }) => {
  const options = {
    fullScreen: {
      enable: true,
      zIndex: -1, // or any value is good for you, if you use -1 set `interactivity.detectsOn` to `"window"` if you need mouse interactions
    },
    // background: {
    //   color: {
    //     value: '#0d47a1',
    //   },
    // },
    fpsLimit: 120,
    interactivity: {
      // events: {
      //   onClick: {
      //     enable: true,
      //     mode: 'push',
      //   },
      //   onHover: {
      //     enable: true,
      //     mode: 'repulse',
      //   },
      // },
      modes: {
        push: {
          quantity: 4,
        },
        repulse: {
          distance: 200,
          duration: 0.4,
        },
      },
    },
    particles: {
      color: {
        value: color,
        opacity: 0.5,
      },
      links: {
        color: color,
        distance: 150,
        enable: true,
        width: 1,
      },
      move: {
        // direction: MoveDirection.none,
        enable: true,
        // outModes: {
        //   default: OutMode.out,
        // },
        random: true,
        speed: 1,
        straight: false,
      },
      number: {
        density: {
          enable: true,
        },
        value: 200,
      },
      opacity: {
        value: 0.5,
      },
      shape: {
        type: 'circle',
      },
      size: {
        value: { min: 1, max: 5 },
      },
    },
    detectRetina: true,
  };

  if (init) {
    return (
      <Particles
        id="tsparticles"
        options={options}
      />
    );
  }

  return <></>;
};

export default ParticlesDemo;
