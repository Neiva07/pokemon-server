import React from 'react';
import Svg, {Path, Circle, LinearGradient, Stop} from 'react-native-svg';

export default {
  Pokeball: (
    <Svg viewBox="0 0 512 512">
      <Path
        d="M256 0C114.844 0 0 114.844 0 256s114.844 256 256 256 256-114.844 256-256S397.156 0 256 0z"
        fill="#303c42"
      />
      <Path
        d="M256 21.333c121.508 0 221.715 92.837 233.483 211.299-17.362 3.845-61.466 11.878-127.904 12.608-5.484-53.704-50.452-95.908-105.579-95.908s-100.095 42.203-105.579 95.908c-66.438-.73-110.542-8.763-127.904-12.608C34.285 114.171 134.492 21.333 256 21.333z"
        fill="#e53935"
      />
      <Circle cx={256} cy={256} r={85.333} fill="#455a64" />
      <Path
        d="M256 490.667C126.604 490.667 21.333 385.396 21.333 256c0-.596.086-1.171.09-1.766 20.31 4.324 64.224 11.547 128.97 12.251 5.357 53.837 50.385 96.181 105.607 96.181s100.25-42.344 105.607-96.181c64.746-.704 108.66-7.927 128.97-12.251.004.595.09 1.169.09 1.766 0 129.396-105.271 234.667-234.667 234.667z"
        fill="#fff"
      />
      <Circle cx={256} cy={256} r={64} fill="#303c42" />
      <Circle cx={256} cy={256} r={42.667} fill="#f2f2f2" />
      <Path
        d="M221.777 268.888c0-23.531 19.135-42.667 42.667-42.667 12.215 0 23.169 5.224 30.954 13.475-6.418-15.456-21.643-26.363-39.398-26.363-23.531 0-42.667 19.135-42.667 42.667 0 11.316 4.501 21.548 11.712 29.191a42.4 42.4 0 01-3.268-16.303z"
        opacity={0.2}
        fill="#fff"
      />
      <LinearGradient
        id="prefix__a"
        gradientUnits="userSpaceOnUse"
        x1={-45.578}
        y1={639.555}
        x2={-23.828}
        y2={629.414}
        gradientTransform="matrix(21.3333 0 0 -21.3333 996.333 13791.667)">
        <Stop offset={0} stopColor="#fff" stopOpacity={0.2} />
        <Stop offset={1} stopColor="#fff" stopOpacity={0} />
      </LinearGradient>
      <Path
        d="M256 0C114.844 0 0 114.844 0 256s114.844 256 256 256 256-114.844 256-256S397.156 0 256 0z"
        fill="url(#prefix__a)"
      />
      <Path
        d="M444.74 239.997c21.049-2.556 36.18-5.467 44.743-7.365-5.801-58.395-33.158-110.496-73.897-148.387C436.051 119.059 448 159.449 448 202.667c0 12.733-1.298 25.142-3.26 37.33zM490.577 254.234c-10.544 2.245-27.586 5.257-50.467 7.747-26.38 100.717-117.906 175.352-226.777 175.352-61.789 0-117.908-24.186-159.859-63.361C94.217 443.637 169.635 490.667 256 490.667c129.396 0 234.667-105.271 234.667-234.667 0-.596-.086-1.171-.09-1.766z"
        opacity={0.1}
      />
    </Svg>
  ),
};
