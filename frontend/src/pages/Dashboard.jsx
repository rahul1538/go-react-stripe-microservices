import { useNavigate } from 'react-router-dom';

const Dashboard = () => {
  const navigate = useNavigate();

  // Function to handle plan selection
  const selectPlan = (planName, price) => {
    // Navigate to the Calculation/Summary page with the plan details
    navigate('/summary', { state: { planName, price } });
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/');
  };

  return (
    <div style={styles.container}>
      <header style={styles.header}>
        <h3>🚀 SaaS Pricing Demo</h3>
        <button onClick={handleLogout} style={styles.logoutBtn}>Logout</button>
      </header>

      <div style={styles.pricingContainer}>
        
        {/* --- Plan 1: Go --- */}
        <div style={styles.card}>
          <h3>Go</h3>
          <div style={styles.price}>
            ₹399 <span style={styles.period}>/ month</span>
          </div>
          <p style={styles.desc}>Do more with smarter AI</p>
          <button 
            style={{...styles.button, background: '#6c5ce7'}} 
            onClick={() => selectPlan('Go Plan', 399)}
          >
            Upgrade to Go
          </button>
          <ul style={styles.features}>
            <li>⚡ Go deep on harder questions</li>
            <li>📷 Chat longer and upload content</li>
            <li>🎨 Make realistic images</li>
          </ul>
        </div>

        {/* --- Plan 2: Plus --- */}
        <div style={styles.card}>
          <h3>Plus</h3>
          <div style={styles.price}>
            ₹1,999 <span style={styles.period}>/ month</span>
          </div>
          <p style={styles.desc}>Unlock the full experience</p>
          <button 
            style={{...styles.button, background: '#000'}} 
            onClick={() => selectPlan('Plus Plan', 1999)}
          >
            Get Plus
          </button>
          <ul style={styles.features}>
            <li>✨ Solve complex problems</li>
            <li>🚀 Create more images, faster</li>
            <li>🧠 Remember goals and context</li>
          </ul>
        </div>

        {/* --- Plan 3: Pro --- */}
        <div style={styles.card}>
          <h3>Pro</h3>
          <div style={styles.price}>
            ₹19,900 <span style={styles.period}>/ month</span>
          </div>
          <p style={styles.desc}>Maximize your productivity</p>
          <button 
            style={{...styles.button, background: '#000'}} 
            onClick={() => selectPlan('Pro Plan', 19900)}
          >
            Get Pro
          </button>
          <ul style={styles.features}>
            <li>🛠️ Master advanced tasks</li>
            <li>📂 Tackle big projects</li>
            <li>⚡ Scale your projects</li>
          </ul>
        </div>

      </div>
    </div>
  );
};

// --- STYLES ---
const styles = {
  container: {
    minHeight: '100vh',
    fontFamily: "'Segoe UI', sans-serif",
    background: '#f9f9f9',
  },
  header: {
    padding: '20px 40px',
    display: 'flex',
    justifyContent: 'space-between',
    alignItems: 'center',
    background: 'white',
    boxShadow: '0 2px 5px rgba(0,0,0,0.05)',
  },
  logoutBtn: {
    padding: '8px 15px',
    background: '#e74c3c',
    color: 'white',
    border: 'none',
    borderRadius: '5px',
    cursor: 'pointer',
  },
  pricingContainer: {
    display: 'flex',
    justifyContent: 'center',
    gap: '20px',
    padding: '50px 20px',
    flexWrap: 'wrap',
  },
  card: {
    background: 'white',
    borderRadius: '15px',
    padding: '30px',
    width: '300px',
    boxShadow: '0 10px 30px rgba(0,0,0,0.05)',
    display: 'flex',
    flexDirection: 'column',
  },
  price: {
    fontSize: '2.5rem',
    fontWeight: 'bold',
    margin: '10px 0',
  },
  period: {
    fontSize: '1rem',
    color: '#666',
    fontWeight: 'normal',
  },
  desc: {
    color: '#666',
    marginBottom: '20px',
  },
  button: {
    padding: '12px',
    borderRadius: '25px',
    color: 'white',
    border: 'none',
    fontWeight: 'bold',
    fontSize: '1rem',
    cursor: 'pointer',
    marginBottom: '20px',
    transition: '0.2s',
  },
  features: {
    listStyle: 'none',
    padding: 0,
    lineHeight: '2rem',
    color: '#444',
  }
};

export default Dashboard;