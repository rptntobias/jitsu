import * as React from "react";

function Svg(props) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      height="100%"
      width="100%"
      viewBox="0 0 200 200"
      {...props}
    >
      <defs>
        <style>
          {
            '.cls-1{fill:#fff;}.cls-2{fill:#ea4335;}.cls-3{fill:#fbbc05;}.cls-4{fill:#4285f4;}.cls-5{fill:#34a853;}.cls-6{fill:#aeaeae;}'
          }
        </style>
      </defs>
      <g id="Guidelines">
        <path
          className="cls-1"
          d="M82.09,30.22a58.69,58.69,0,0,0-33.1,29A57.63,57.63,0,0,0,43.9,73.88a58.44,58.44,0,0,0,42.14,68,62.48,62.48,0,0,0,29.9.31,51.54,51.54,0,0,0,38.72-36.66,68.55,68.55,0,0,0,1.48-31h-55V97.26H133a27.51,27.51,0,0,1-11.69,18,33.44,33.44,0,0,1-12.87,5.08,38.48,38.48,0,0,1-14,0,34.61,34.61,0,0,1-13-5.62A36.08,36.08,0,0,1,68.14,96.82a35.42,35.42,0,0,1,0-22.72,36.47,36.47,0,0,1,8.48-13.78,34.36,34.36,0,0,1,34.61-9,31.36,31.36,0,0,1,12.59,7.41L134.56,48c1.89-1.91,3.87-3.77,5.69-5.74a57.1,57.1,0,0,0-18.81-11.65A59,59,0,0,0,82.09,30.22Z"
        />
        <path
          className="cls-2"
          d="M82.09,30.22a59.16,59.16,0,0,1,39.35.34,56.68,56.68,0,0,1,18.81,11.65c-1.82,2-3.8,3.83-5.69,5.74L123.82,58.68a31.69,31.69,0,0,0-12.59-7.4,34.37,34.37,0,0,0-17.15-.69,34.9,34.9,0,0,0-17.46,9.7,36.27,36.27,0,0,0-8.48,13.77c-6.37-5-12.75-9.88-19.15-14.84A58.63,58.63,0,0,1,82.09,30.22Z"
        />
        <path
          className="cls-3"
          d="M43.93,73.84A58.1,58.1,0,0,1,49,59.16Q58.58,66.59,68.18,74a35.12,35.12,0,0,0,0,22.73q-9.57,7.44-19.13,14.84A58.12,58.12,0,0,1,43.93,73.84Z"
        />
        <path
          className="cls-4"
          d="M101.18,74.44h55a69,69,0,0,1-1.48,31,53,53,0,0,1-14.79,24.23l-18.56-14.4a27.48,27.48,0,0,0,11.68-18H101.15Q101.2,85.83,101.18,74.44Z"
        />
        <path
          className="cls-5"
          d="M49,111.6q9.56-7.38,19.12-14.84a36.25,36.25,0,0,0,13.38,17.92,34.82,34.82,0,0,0,13,5.62,38,38,0,0,0,14,0,33.46,33.46,0,0,0,12.87-5.09l18.56,14.41a52.73,52.73,0,0,1-23.93,12.43,62.61,62.61,0,0,1-29.9-.31,57.68,57.68,0,0,1-21.22-10.71A58.64,58.64,0,0,1,49,111.6Z"
        />
        <path
          className="cls-6"
          d="M40.55,160.2l-1.8,1.06a3.21,3.21,0,0,0-1-1.14,2.6,2.6,0,0,0-2.77.22,1.66,1.66,0,0,0-.61,1.3c0,.71.54,1.29,1.6,1.73l1.47.6a6.21,6.21,0,0,1,2.62,1.77,4,4,0,0,1,.83,2.55,4.51,4.51,0,0,1-1.35,3.36A4.65,4.65,0,0,1,36.19,173,4.55,4.55,0,0,1,33,171.85a5.21,5.21,0,0,1-1.53-3.2l2.25-.49a3.75,3.75,0,0,0,.53,1.79,2.54,2.54,0,0,0,3.72.26,2.4,2.4,0,0,0,.69-1.76,2.47,2.47,0,0,0-.12-.79,2,2,0,0,0-.37-.66,3.33,3.33,0,0,0-.65-.56,7.08,7.08,0,0,0-1-.5l-1.42-.59q-3-1.28-3-3.73a3.54,3.54,0,0,1,1.27-2.77,4.6,4.6,0,0,1,3.16-1.13A4.41,4.41,0,0,1,40.55,160.2Z"
        />
        <path
          className="cls-6"
          d="M51.65,168.58H45a2.82,2.82,0,0,0,.74,1.82,2.29,2.29,0,0,0,1.69.67,2.15,2.15,0,0,0,1.32-.39,5.19,5.19,0,0,0,1.18-1.41l1.81,1a7.54,7.54,0,0,1-.89,1.22,4.6,4.6,0,0,1-1,.84,4,4,0,0,1-1.15.48,5.76,5.76,0,0,1-1.35.15A4.36,4.36,0,0,1,44,171.65a5,5,0,0,1-1.26-3.57A5.16,5.16,0,0,1,44,164.52a4.23,4.23,0,0,1,3.26-1.34,4.18,4.18,0,0,1,3.24,1.3,5.14,5.14,0,0,1,1.18,3.6Zm-2.2-1.75a2.05,2.05,0,0,0-2.16-1.72,2.25,2.25,0,0,0-.74.12,2.1,2.1,0,0,0-1.1.88,2.34,2.34,0,0,0-.31.72Z"
        />
        <path
          className="cls-6"
          d="M60.84,163.44H63v9.27H60.84v-1a4.13,4.13,0,0,1-6-.15,5.18,5.18,0,0,1-1.24-3.54,5,5,0,0,1,1.24-3.48,4,4,0,0,1,3.13-1.39,4,4,0,0,1,2.91,1.33Zm-5.09,4.61a3.22,3.22,0,0,0,.71,2.17,2.37,2.37,0,0,0,1.85.85,2.5,2.5,0,0,0,1.93-.82,3.55,3.55,0,0,0,0-4.31,2.45,2.45,0,0,0-1.91-.83,2.37,2.37,0,0,0-1.85.84A3.09,3.09,0,0,0,55.75,168.05Z"
        />
        <path
          className="cls-6"
          d="M65.7,163.44h2.14v.83a3.78,3.78,0,0,1,1-.85,2.38,2.38,0,0,1,1.1-.24,3.47,3.47,0,0,1,1.78.55l-1,2a2,2,0,0,0-1.19-.43c-1.17,0-1.75.88-1.75,2.64v4.81H65.7Z"
        />
        <path
          className="cls-6"
          d="M80.05,163.84v2.84a5.5,5.5,0,0,0-1.32-1.23,2.8,2.8,0,0,0-3.35.51,2.94,2.94,0,0,0-.8,2.12,3,3,0,0,0,.77,2.15,2.59,2.59,0,0,0,2,.84,2.74,2.74,0,0,0,1.36-.34,5.39,5.39,0,0,0,1.35-1.25v2.82a5.64,5.64,0,0,1-2.61.68,5,5,0,0,1-3.61-1.39,4.93,4.93,0,0,1,0-7,5,5,0,0,1,3.6-1.43A5.4,5.4,0,0,1,80.05,163.84Z"
        />
        <path
          className="cls-6"
          d="M82.63,156.63h2.14v7.58a3.7,3.7,0,0,1,2.53-1,3.23,3.23,0,0,1,2.51,1,4.11,4.11,0,0,1,.81,2.83v5.68H88.47v-5.48a2.64,2.64,0,0,0-.39-1.62,1.55,1.55,0,0,0-1.28-.5,1.76,1.76,0,0,0-1.58.7,4.76,4.76,0,0,0-.45,2.42v4.48H82.63Z"
        />
        <path
          className="cls-6"
          d="M110.64,158.8v2.63a6.16,6.16,0,0,0-4-1.61,5.07,5.07,0,0,0-3.82,1.63,5.49,5.49,0,0,0-1.56,4,5.36,5.36,0,0,0,1.56,3.9,5.16,5.16,0,0,0,3.83,1.59,4.59,4.59,0,0,0,2-.38,6.35,6.35,0,0,0,.95-.51c.33-.22.67-.48,1-.78v2.67a8,8,0,0,1-4,1.08,7.37,7.37,0,0,1-5.38-2.2A7.29,7.29,0,0,1,99,165.43a7.54,7.54,0,0,1,1.86-5,7.37,7.37,0,0,1,5.91-2.7A7.52,7.52,0,0,1,110.64,158.8Z"
        />
        <path
          className="cls-6"
          d="M112.59,168a4.61,4.61,0,0,1,1.44-3.42,5.07,5.07,0,0,1,7,0,5,5,0,0,1,0,7,4.94,4.94,0,0,1-3.56,1.4,4.71,4.71,0,0,1-3.49-1.43A4.8,4.8,0,0,1,112.59,168Zm2.19,0a3.16,3.16,0,0,0,.74,2.2,2.64,2.64,0,0,0,2,.82,2.58,2.58,0,0,0,2-.82,3,3,0,0,0,.76-2.16,3,3,0,0,0-.76-2.16,2.85,2.85,0,0,0-4,0A3,3,0,0,0,114.78,168.05Z"
        />
        <path
          className="cls-6"
          d="M124.76,163.44h2.15v.85a3.51,3.51,0,0,1,2.53-1.11,3.21,3.21,0,0,1,2.53,1,4.18,4.18,0,0,1,.78,2.83v5.68H130.6v-5.18a3.43,3.43,0,0,0-.38-1.89,1.54,1.54,0,0,0-1.36-.54,1.65,1.65,0,0,0-1.51.71,4.85,4.85,0,0,0-.44,2.43v4.47h-2.15Z"
        />
        <path
          className="cls-6"
          d="M140.85,165l-1.77.94c-.28-.57-.63-.86-1-.86a.73.73,0,0,0-.51.2.64.64,0,0,0-.21.5c0,.35.42.71,1.24,1.06a7.22,7.22,0,0,1,2.3,1.35,2.32,2.32,0,0,1,.59,1.66,2.91,2.91,0,0,1-1,2.25,3.55,3.55,0,0,1-5.63-1.42l1.83-.84a3.61,3.61,0,0,0,.58.84,1.29,1.29,0,0,0,.93.37c.73,0,1.09-.34,1.09-1,0-.38-.28-.73-.84-1.06l-.65-.32-.66-.31a4,4,0,0,1-1.31-.91,2.29,2.29,0,0,1-.49-1.5,2.65,2.65,0,0,1,.83-2,2.9,2.9,0,0,1,2.06-.79A2.8,2.8,0,0,1,140.85,165Z"
        />
        <path
          className="cls-6"
          d="M142.94,168a4.64,4.64,0,0,1,1.43-3.42,5.09,5.09,0,0,1,7,0,5,5,0,0,1,0,7,4.94,4.94,0,0,1-3.56,1.4,4.71,4.71,0,0,1-3.49-1.43A4.8,4.8,0,0,1,142.94,168Zm2.19,0a3.16,3.16,0,0,0,.74,2.2,2.64,2.64,0,0,0,2,.82,2.58,2.58,0,0,0,2-.82,3,3,0,0,0,.76-2.16,3.06,3.06,0,0,0-.76-2.16,2.64,2.64,0,0,0-2-.82,2.62,2.62,0,0,0-2,.82A3,3,0,0,0,145.13,168.05Z"
        />
        <path className="cls-6" d="M157.25,156.63v16.08H155.1V156.63Z" />
        <path
          className="cls-6"
          d="M168.42,168.58h-6.65a2.77,2.77,0,0,0,.74,1.82,2.29,2.29,0,0,0,1.69.67,2.15,2.15,0,0,0,1.32-.39,5.38,5.38,0,0,0,1.17-1.41l1.81,1a7.52,7.52,0,0,1-.88,1.22,4.85,4.85,0,0,1-1,.84,4,4,0,0,1-1.16.48,5.67,5.67,0,0,1-1.34.15,4.35,4.35,0,0,1-3.33-1.33,5,5,0,0,1-1.26-3.57,5.12,5.12,0,0,1,1.22-3.56,4.66,4.66,0,0,1,6.5,0,5.19,5.19,0,0,1,1.18,3.6Zm-2.2-1.75a2.06,2.06,0,0,0-2.17-1.72,2.19,2.19,0,0,0-.73.12,2,2,0,0,0-.62.34,2.31,2.31,0,0,0-.49.54,2.6,2.6,0,0,0-.3.72Z"
        />
      </g>
    </svg>
  );
}

export default Svg;