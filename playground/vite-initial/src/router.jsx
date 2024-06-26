import { useEffect, useState } from 'react';
import { createHashRouter, Link, Outlet } from 'react-router-dom';
import { getFingerprint } from '@thumbmarkjs/thumbmarkjs';

import { initParticlesEngine } from '@tsparticles/react';
import { loadSlim } from '@tsparticles/slim'; // if you are going to use `loadSlim`, install the "@tsparticles/slim" package too.

import { API_URL } from './config';
import Tilt from 'react-parallax-tilt';
import './tilt.css';
import Device from './Device';
import ParticlesDemo from './ParticlesDemo';

const Home = () => {
  return (
    <div>
      <h2>tablechat.me</h2>
      <div>A new way to chat</div>
    </div>
  );
};

const Fingerprint = () => {
  const [fingerprint, setFingerprint] = useState('');

  useEffect(() => {
    getFingerprint()
      .then((result) => {
        setFingerprint(result);
      })
      .catch((error) => {
        console.error('Error getting fingerprint:', error);
      });
  }, []);

  return (
    <div>
      <h2>Fingerprint</h2>
      <div>{fingerprint}</div>
    </div>
  );
};

const Phone = () => {
  const [color, setColor] = useState('#0d47a1');
  const [init, setInit] = useState(false);

  // this should be run only once per application lifetime
  useEffect(() => {
    initParticlesEngine(async (engine) => {
      // you can initiate the tsParticles instance (engine) here, adding custom shapes or presets
      // this loads the tsparticles package bundle, it's the easiest method for getting everything ready
      // starting from v2 you can add only the features you need reducing the bundle size
      //await loadAll(engine);
      //await loadFull(engine);
      await loadSlim(engine);
      //await loadBasic(engine);
    }).then(() => {
      setInit(true);
    });
  }, []);

  const handleClick = () => {
    setColor('#' + Math.floor(Math.random() * 16777215).toString(16));
  }

  return (
    <div>
      <div
        style={{
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
        }}
        onClick={handleClick}
      >
        <Tilt
          glareEnable={true}
          glareMaxOpacity={0.8}
          glareColor="#ffffff"
          glarePosition="bottom"
          glareBorderRadius="20px"
          className="default-component"
        >
          <Device />
        </Tilt>
      </div>
      <ParticlesDemo color={color} init={init} key={color} />
      <input
        type="color"
        value={color}
        onChange={(e) => setColor(e.target.value)}
      />
    </div>
  );
};

const About = () => {
  return (
    <div>
      <h2>About</h2>
      <div>Networking and entertainment</div>
      <div>powered by cloudflare</div>
    </div>
  );
};

const Healthz = () => {
  const [status, setStatus] = useState('Loading...');

  useEffect(() => {
    fetch(`${API_URL}/api/healthz`)
      .then((response) => response.json())
      .then((data) => setStatus(data.status));
  }, []);

  return (
    <div>
      <h2>Healthz</h2>
      <div>{status}</div>
    </div>
  );
};

const PrivacyPolicy = () => {
  return (
    <div>
      <h2>Privacy Policy</h2>
      <div>
        <h3>Information Collection and Use</h3>
        We collect users' email addresses solely for the purpose of
        authenticating them to our service through Google OAuth. We do not use
        this information for any other purposes.
        <h3>Information Sharing</h3>
        We do not share users' email addresses with any third parties. Your
        email address is only used for authentication purposes within our
        application.
        <h3>Data Security</h3>
        We take appropriate measures to protect the security of users' email
        addresses. However, please be aware that no method of transmission over
        the internet, or method of electronic storage, is 100% secure.
        Therefore, while we strive to use commercially acceptable means to
        protect your email address, we cannot guarantee its absolute security.
        <h3>Access to Personal Information</h3>
        Users can access and update their email address associated with our
        service through the Google account settings. We do not store any
        additional personal information beyond the email address provided during
        authentication.
        <h3>Changes to Privacy Policy</h3>
        We reserve the right to update our Privacy Policy from time to time. Any
        changes will be posted on this page, and users will be notified via
        email.
        <h3>Contact Us</h3>
        If you have any questions or concerns about our Privacy Policy, please
        contact us at [hello at tablechat.me].
      </div>
    </div>
  );
};

const TermsOfService = () => {
  return (
    <div>
      <h2>Terms Of Service</h2>
      <div>
        <h3>Acceptance of Terms</h3>
        By accessing or using our service, you agree to be bound by these Terms
        of Service. If you do not agree with any part of these terms, you may
        not access the service.
        <h3>User Responsibilities</h3>
        You are solely responsible for maintaining the confidentiality of your
        account credentials. You agree to notify us immediately of any
        unauthorized use of your account.
        <h3>Limitation of Liability</h3>
        We are not responsible for any content posted, uploaded, or otherwise
        made available by users of our service. You use our service at your own
        risk.
        <h3>Indemnification</h3>
        You agree to indemnify and hold us harmless from any claims, damages,
        losses, or liabilities arising out of your use of our service or
        violation of these Terms of Service.
        <h3>Changes to Terms of Service</h3>
        We reserve the right to update or modify these Terms of Service at any
        time without prior notice. Any changes will be effective immediately
        upon posting on this page.
        <h3>Governing Law</h3>
        These Terms of Service shall be governed by and construed in accordance
        with the law, without regard to its conflict of law provisions.
        <h3>Contact Us</h3>
        If you have any questions or concerns about our Terms of Service, please
        contact us at [hello at tablechat.me].
      </div>
    </div>
  );
};

const FAQ = () => {
  return (
    <div>
      <h2>FAQ</h2>
      <div>Stay tuned for more</div>
    </div>
  );
};

const Root = () => {
  return (
    <div>
      <h1>tablechat.me</h1>
      <div>
        <ul>
          <li>
            <Link to="/">Home</Link>
          </li>
          <li>
            <Link to="/about">About</Link>
          </li>
          <li>
            <Link to="/privacy-policy">Privacy Policy</Link>
          </li>
          <li>
            <Link to="/terms-of-service">Terms Of Service</Link>
          </li>
          <li>
            <Link to="/faq">FAQ</Link>
          </li>
          <li>
            <Link to="/healthz">Healthz</Link>
          </li>
          <li>
            <Link to="/phone">Phone</Link>
          </li>
          <li>
            <Link to="/fingerprint">Fingerprint</Link>
          </li>
        </ul>
      </div>
      <Outlet />
    </div>
  );
};

const routes = [
  {
    path: '/',
    element: <Root />,
    children: [
      {
        path: '/',
        element: <Home />,
      },
      {
        path: '/about',
        element: <About />,
      },
      {
        path: '/privacy-policy',
        element: <PrivacyPolicy />,
      },
      {
        path: '/terms-of-service',
        element: <TermsOfService />,
      },
      {
        path: '/faq',
        element: <FAQ />,
      },
      {
        path: '/healthz',
        element: <Healthz />,
      },
      {
        path: '/phone',
        element: <Phone />,
      },
      {
        path: '/fingerprint',
        element: <Fingerprint />,
      },
    ],
  },
];

const router = createHashRouter(routes);

export default router;
