import { Component } from 'react';
import ErrorFallback from './ErrorFallback';

export default class ErrorBoundary extends Component {
  constructor(props) {
    super(props);
    this.state = { hasError: false };
  }

  static getDerivedStateFromError(error) {
    return ({ hasError: true, error });
  }

  componentDidCatch(error, info) {
    this.state = { ...this.state, hasError: true };
  }

  render() {
    const { hasError, error } = this.state;
    if (hasError) {
      return (
        <ErrorFallback errorName={error.name} info={error.message} />
      );
    }
    const { children } = this.props;
    return children;
  }
}
