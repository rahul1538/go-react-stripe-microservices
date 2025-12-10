import { useState } from 'react';
import { createCheckoutSession } from '../api/api';
import { useNavigate } from 'react-router-dom';

const Checkout = () => {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);

  const handlePayment = async () => {
    setLoading(true);
    try {
      // 1. Ask Backend for the Stripe URL
      const { data } = await createCheckoutSession({ amount: 1000, currency: 'usd' });
      
      // 2. Redirect the user to Stripe's Hosted Page
      if (data.url) {
        window.location.href = data.url;
      } else {
        alert("Failed to get payment URL");
      }
    } catch (err) {
      console.error(err);
      alert('Payment initialization failed.');
      setLoading(false);
    }
  };

  return (
    <div style={styles.container}>
      <div style={styles.card}>
        <div style={styles.header}>
          <h2>Stripe Hosted Checkout</h2>
          <p>Secure Payment Pipeline</p>
        </div>
        
        <div style={styles.summary}>
          <span>Invoice Total:</span>
          <span style={styles.amount}>$10.00</span>
        </div>

        <p style={{textAlign: 'center', color: '#666', marginBottom: '20px'}}>
          You will be redirected to the secure Stripe payment page to complete your transaction.
        </p>

        <button 
          onClick={handlePayment} 
          disabled={loading} 
          style={styles.button}
        >
          {loading ? 'Redirecting to Stripe...' : 'Proceed to Pay $10.00'}
        </button>

        <button onClick={() => navigate('/dashboard')} style={styles.cancelBtn}>
          Cancel
        </button>
      </div>
    </div>
  );
};

const styles = {
  container: {
    minHeight: '100vh',
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    background: '#f7f9fc',
    fontFamily: "'Segoe UI', sans-serif",
  },
  card: {
    background: 'white',
    padding: '40px',
    borderRadius: '15px',
    boxShadow: '0 15px 35px rgba(0,0,0,0.1)',
    width: '100%',
    maxWidth: '450px',
  },
  header: {
    textAlign: 'center',
    marginBottom: '30px',
  },
  summary: {
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    padding: '15px',
    background: '#f8f9fa',
    borderRadius: '10px',
    marginBottom: '20px',
    border: '1px solid #e9ecef',
  },
  amount: {
    fontSize: '1.5rem',
    fontWeight: 'bold',
    color: '#333',
  },
  button: {
    width: '100%',
    padding: '15px',
    background: '#635bff', // Stripe Brand Purple
    color: 'white',
    border: 'none',
    borderRadius: '25px',
    fontSize: '1.1rem',
    fontWeight: 'bold',
    cursor: 'pointer',
    transition: '0.2s',
    boxShadow: '0 4px 6px rgba(50, 50, 93, 0.11), 0 1px 3px rgba(0, 0, 0, 0.08)',
  },
  cancelBtn: {
    marginTop: '15px',
    background: 'none',
    border: 'none',
    color: '#888',
    cursor: 'pointer',
    width: '100%',
  }
};

export default Checkout;