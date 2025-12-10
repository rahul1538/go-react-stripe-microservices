import { useLocation, useNavigate } from 'react-router-dom';
import { useState } from 'react';
import { createCheckoutSession } from '../api/api';

const OrderSummary = () => {
  const { state } = useLocation(); // Get data passed from Dashboard
  const navigate = useNavigate();
  const [loading, setLoading] = useState(false);

  // If someone visits directly without selecting a plan, go back
  if (!state) {
    return <div style={{padding: '50px'}}>Please select a plan from Dashboard.</div>;
  }

  const { planName, price } = state;
  const taxRate = 0.18; // 18% GST
  const taxAmount = price * taxRate;
  const totalAmount = price + taxAmount;

  const handleProceedToPay = async () => {
    setLoading(true);
    try {
      // Convert to cents/paise for Stripe (Total * 100)
      const amountInSmallestUnit = Math.round(totalAmount * 100);

      // Call Backend to get Stripe URL
      const { data } = await createCheckoutSession({ 
        amount: amountInSmallestUnit, 
        currency: 'inr' // Using INR since your screenshot showed ₹
      });

      // Redirect to Stripe
      if (data.url) {
        window.location.href = data.url;
      }
    } catch (err) {
      console.error(err);
      alert('Payment failed to start.');
      setLoading(false);
    }
  };

  return (
    <div style={styles.container}>
      <div style={styles.summaryCard}>
        <h2>🧾 Order Summary</h2>
        <hr style={{border: '0', borderTop: '1px solid #eee', margin: '20px 0'}}/>
        
        <div style={styles.row}>
          <span>Plan Selected:</span>
          <strong>{planName}</strong>
        </div>

        <div style={styles.row}>
          <span>Base Price:</span>
          <span>₹{price.toLocaleString()}</span>
        </div>

        <div style={styles.row}>
          <span>GST (18%):</span>
          <span>₹{taxAmount.toLocaleString()}</span>
        </div>

        <hr style={{border: '0', borderTop: '1px dashed #ccc', margin: '20px 0'}}/>

        <div style={{...styles.row, fontSize: '1.4rem', fontWeight: 'bold'}}>
          <span>Total:</span>
          <span style={{color: '#27ae60'}}>₹{totalAmount.toLocaleString()}</span>
        </div>

        <button 
          onClick={handleProceedToPay} 
          style={styles.payBtn}
          disabled={loading}
        >
          {loading ? 'Redirecting to Stripe...' : `Pay ₹${totalAmount.toLocaleString()}`}
        </button>

        <button onClick={() => navigate('/dashboard')} style={styles.cancelBtn}>
          Cancel / Back
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
    background: '#f4f6f8',
    fontFamily: "'Segoe UI', sans-serif",
  },
  summaryCard: {
    background: 'white',
    padding: '40px',
    borderRadius: '15px',
    width: '100%',
    maxWidth: '500px',
    boxShadow: '0 10px 40px rgba(0,0,0,0.1)',
  },
  row: {
    display: 'flex',
    justifyContent: 'space-between',
    marginBottom: '15px',
    fontSize: '1.1rem',
    color: '#333',
  },
  payBtn: {
    width: '100%',
    padding: '15px',
    background: '#27ae60',
    color: 'white',
    border: 'none',
    borderRadius: '10px',
    fontSize: '1.2rem',
    fontWeight: 'bold',
    cursor: 'pointer',
    marginTop: '20px',
    transition: '0.2s',
  },
  cancelBtn: {
    width: '100%',
    padding: '10px',
    background: 'transparent',
    border: 'none',
    color: '#777',
    cursor: 'pointer',
    marginTop: '10px',
  }
};

export default OrderSummary;