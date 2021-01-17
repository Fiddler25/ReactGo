import React, { Component } from "react";
import LoginPage from "./LoginPage";
import HomePage from "./HomePage";

export default class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      hasLoggedIn: false,
      userId: "",
      passwourd: "",
    };
  }

  onLogin = async (userId, password) => {
    const endpoint = "//localhost:8080/login";
    const options = {
      method: "POST",
      headers: {
        "Content-Type": "application/json; charset=utf-8",
      },
      body: JSON.stringify({ userId, password }),
    };
    const { result } = await fetch(endpoint, options).then((res) => res.json());
    console.log(result);
    this.setState({ hasLoggedIn: result });
  };

  onLogout = () => {
    this.setState({ hasLoggedIn: false });
  };

  handleChangeText = (userId) => {
    this.setState({ userId });
  };

  handleChangePassword = (password) => {
    this.setState({ password });
  };

  render() {
    const { userId, password } = this.state;
    return (
      <div
        style={{
          width: "100vw",
          height: "100vh",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        {this.state.hasLoggedIn ? (
          <HomePage onLogout={this.onLogout} userId={userId} />
        ) : (
          <LoginPage
            onLogin={this.onLogin}
            onChangeText={this.handleChangeText}
            onChangePassword={this.handleChangePassword}
            userId={userId}
            password={password}
          />
        )}
      </div>
    );
  }
}
