import React, { Component } from 'react';
import ReactDOM from "react-dom";
import Style from './{{.Stylesheets}}/{{.DefaultCss}}';

class App extends Component {

  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div>Hello, {{.Name}}</div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById("root"));
