var ReactDOM = require("react-dom");
import React from 'react';
// import { Page } from './widgets/page';
const e = React.createElement;

const root = document.getElementById('page_root');

ReactDOM.unstable_createRoot(root).render(e(Counter));