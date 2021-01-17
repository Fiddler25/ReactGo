import React, { Component } from "react";

export default class HomePage extends Component {
  render() {
    return (
      <div align="center">
        <p>ようこそ {"test"} さん</p>
        <button onClick={this.props.onLogout}>ログアウト</button>
      </div>
    );
  }
}
