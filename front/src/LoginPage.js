import React, { Component } from "react";

export default class LoginPage extends Component {
  handleChangeText = (event) => {
    this.props.onChangeText(event.target.value);
  };

  handleChangePassword = (event) => {
    this.props.onChangePassword(event.target.value);
  };

  onLogin = () => {
    const { userId, password } = this.props;
    this.props.onLogin(userId, password);
  };

  render() {
    const { userId, password } = this.props;
    return (
      <div align="center">
        <table border="0">
          <tbody>
            <tr>
              <th>ユーザID</th>
              <td>
                <input
                  type="text"
                  label="User ID"
                  maxLength="10"
                  value={userId}
                  onChange={this.handleChangeText}
                />
              </td>
            </tr>
            <tr>
              <th>パスワード</th>
              <td>
                <input
                  type="password"
                  label="password"
                  maxLength="10"
                  onChange={this.handleChangePassword}
                />
              </td>
            </tr>
          </tbody>
        </table>
        <button onClick={this.onLogin}>ログイン</button>
      </div>
    );
  }
}
