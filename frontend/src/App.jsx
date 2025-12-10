import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import OrderSummary from './pages/OrderSummary'; // Import the new page
import Checkout from './pages/Checkout'; // Keep your old checkout as backup if needed

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/dashboard" element={<Dashboard />} />
        
        {/* The new Calculation Page */}
        <Route path="/summary" element={<OrderSummary />} />
        
        {/* Optional: Keep this if you want to access the old demo page */}
        <Route path="/checkout" element={<Checkout />} />
      </Routes>
    </Router>
  );
}

export default App;