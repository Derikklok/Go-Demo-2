# Stress Detection System - Backend API Documentation

## 🎯 What is this project?

This is a **backend API server** for a stress detection system that analyzes eye images to determine if a person is experiencing stress. The system examines specific eye characteristics like **pupil dilation** and **tension rings** in the iris to make this determination.

Think of it like a digital doctor that can look at your eyes and tell you if you're stressed - similar to how a doctor might notice if your pupils are dilated or if there are signs of strain in your eyes.

## 🔍 How does it work?

### The Big Picture
1. **Users register** with their username, password, and birthdate
2. **Users upload eye images** through the system
3. **The system analyzes** these images for stress indicators
4. **Results are stored** and can be retrieved later
5. **AI recommendations** are provided based on analysis history

### Key Eye Analysis Features
- **Pupil Dilation Measurement**: How wide or narrow your pupils are
- **Tension Ring Detection**: Circular patterns in the iris that may indicate stress
- **Stress Classification**: Boolean result (true/false) if stress is detected
- **Confidence Level**: How confident the AI is about its analysis

## 🏗️ Project Architecture

### Core Components

```
📁 Project Structure
├── 🔧 configs/           # Configuration files
│   ├── db.config.js      # Database connection setup
│   ├── cloudinary.config.js # Image storage setup
│   └── gemini.config.js  # AI assistant setup
│
├── 🎮 controllers/       # Business logic handlers
│   ├── auth.controller.js        # User login/registration
│   ├── analysis.controller.js    # Stress analysis operations
│   ├── upload.controller.js      # Image upload handling
│   └── recommendations.controller.js # AI recommendations
│
├── 🛡️ middlewares/       # Request processing
│   └── upload.middleware.js # File upload validation
│
├── 🛣️ routes/            # API endpoint definitions
│   ├── auth.routes.js     # /api/auth/* endpoints
│   ├── analysis.routes.js # /api/analysis/* endpoints
│   ├── upload.routes.js   # /api/upload/* endpoints
│   └── recommendations.routes.js # /api/recommendations/* endpoints
│
├── 📁 uploads/           # Temporary file storage
├── 🚀 index.js           # Main server file
└── 📦 package.json       # Dependencies and scripts
```

### Technology Stack

- **Runtime**: Node.js (JavaScript server environment)
- **Framework**: Express.js (web application framework)
- **Database**: MongoDB (document-based database)
- **Image Storage**: Cloudinary (cloud-based image management)
- **AI Assistant**: Google Gemini (for generating recommendations)
- **File Upload**: Multer (handling multipart/form-data)

## 📚 Database Collections

The system uses MongoDB with the following collections:

### 1. Users Collection
```javascript
{
  _id: ObjectId,
  username: "john_doe",
  password: "user_password", // Note: Not hashed in current version
  birthDate: ISODate("1990-05-15"),
  createdAt: ISODate("2025-11-20T10:30:00.000Z")
}
```

### 2. Analyses Collection
```javascript
{
  _id: ObjectId,
  username: "john_doe",
  hasStress: true,              // Main result: stressed or not
  imageUrl: "https://cloudinary.../image.jpg",
  confidenceLevel: 0.85,        // AI confidence (0-1)
  pupilDilation: 4.2,          // Pupil size measurement
  tensionRings: 3,             // Number of tension rings detected
  fullDetails: {               // Complete analysis from AI model
    // Additional technical details
  },
  createdAt: ISODate("2025-11-20T10:30:00.000Z")
}
```

### 3. Eye Images Collection
```javascript
{
  _id: ObjectId,
  username: "john_doe",
  imageUrl: "https://cloudinary.../image.jpg",
  cloudinaryId: "eye_glaze_images/image-123456",
  uploadedAt: ISODate("2025-11-20T10:30:00.000Z")
}
```

## 🌐 API Endpoints Guide

### Authentication System (`/api/auth`)

#### Register New User
**POST** `/api/auth/register`
```javascript
// What you send:
{
  "username": "john_doe",
  "password": "secure_password",
  "birthDate": "1990-05-15"  // Required format: YYYY-MM-DD
}

// What you get back:
{
  "status": "success",
  "message": "User registered successfully",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "username": "john_doe",
    "age": 35,  // Automatically calculated from birthDate
    "createdAt": "2025-11-20T10:30:00.000Z"
  }
}
```

#### Login User
**POST** `/api/auth/login`
```javascript
// What you send:
{
  "username": "john_doe",
  "password": "secure_password"
}

// What you get back:
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "username": "john_doe",
    "age": 35,
    "createdAt": "2025-11-20T10:30:00.000Z"
  }
}
```

### Image Upload System (`/api/upload`)

#### Upload Eye Image
**POST** `/api/upload/eye-image`
```javascript
// Form data (multipart/form-data):
username: "john_doe"
image: [eye_image_file.jpg]

// What you get back:
{
  "status": "success",
  "message": "Eye image uploaded successfully",
  "data": {
    "id": "655e7d2fc9e5c123456789ef",
    "username": "john_doe",
    "imageUrl": "https://res.cloudinary.com/yourcloud/image.jpg",
    "uploadedAt": "2025-11-20T10:30:00.000Z"
  }
}
```

#### Get User's Images
**GET** `/api/upload/eye-images/:username`
```javascript
// Returns array of all images for the user, newest first
```

### Stress Analysis System (`/api/analysis`)

#### Submit Analysis Result
**POST** `/api/analysis/submit`
```javascript
// What you send (minimum required):
{
  "username": "john_doe",
  "hasStress": true,
  "pupilDilation": 4.2,
  "tensionRings": 3
}

// What you send (complete):
{
  "username": "john_doe",
  "hasStress": true,
  "imageUrl": "https://cloudinary.../image.jpg",
  "confidenceLevel": 0.85,
  "pupilDilation": 4.2,
  "tensionRings": 3,
  "fullDetails": {
    "iris_analysis": {
      "tension_rings_detected": 3,
      "color_variations": ["brown", "amber"]
    },
    "pupil_analysis": {
      "dilation_level": 4.2,
      "responsiveness": "normal"
    }
  }
}

// What you get back:
{
  "status": "success",
  "message": "Analysis submitted successfully",
  "data": {
    "id": "655d4b3fc9e5c123456789ab",
    "username": "john_doe",
    "hasStress": true,
    "imageUrl": "https://cloudinary.../image.jpg",
    "pupilDilation": 4.2,
    "tensionRings": 3,
    "createdAt": "2025-11-20T10:30:00.000Z"
  }
}
```

#### Get User's Analysis History
**GET** `/api/analysis/user/:username`
- Returns all analyses for a user, sorted by newest first

#### Get Latest Analysis
**GET** `/api/analysis/latest/:username`
- Returns the most recent analysis for a user

#### Get Analysis Count
**GET** `/api/analysis/count/:username`
- Returns total number of analyses performed for a user

### AI Recommendations (`/api/recommendations`)

#### Get Personalized Recommendations
**GET** `/api/recommendations/user/:username`
```javascript
// What you get back:
{
  "status": "success",
  "message": "Recommendations generated successfully",
  "data": {
    "analysis": {
      "hasStress": true,
      "createdAt": "2025-11-20T10:30:00.000Z"
    },
    "stats": {
      "totalAnalysesLastWeek": 5,
      "stressDetectedCount": 3,
      "stressPercentage": 60
    },
    "recommendations": {
      "assessment": "Based on your recent eye analyses, you're experiencing moderate eye strain...",
      "recommendations": [
        "Take regular 20-20-20 breaks",
        "Adjust screen brightness to match surroundings",
        "Consider blue light filtering glasses"
      ],
      "lifestyleAdjustments": [
        "Stay hydrated throughout the day",
        "Ensure proper lighting in workspace"
      ]
    }
  }
}
```

## 🔧 Environment Configuration

Create a `.env` file in your project root with:

```bash
# Server Configuration
EXPRESS_PORT=8000

# Database Configuration
EXPRESS_MONGO_NODE_SRV=mongodb+srv://username:password@cluster.mongodb.net/

# Cloudinary Configuration (for image storage)
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret

# Google Gemini AI Configuration (for recommendations)
GEMINI_API_KEY=your_gemini_api_key
```

## 🚀 Getting Started

### Prerequisites
- **Node.js** (version 16 or higher)
- **MongoDB Atlas** account (or local MongoDB installation)
- **Cloudinary** account (for image storage)
- **Google Gemini API** key (for AI recommendations)

### Installation Steps

1. **Clone and Setup**
   ```bash
   cd your-project-directory
   npm install
   ```

2. **Configure Environment**
   - Create `.env` file with the configuration above
   - Replace placeholder values with your actual credentials

3. **Run the Server**
   ```bash
   # Development mode (with auto-restart)
   npm run dev

   # Production mode
   npm start
   ```

4. **Verify Installation**
   - Open browser to `http://localhost:8000`
   - You should see: `{"status": "API is running"}`
   - Check database connection: `http://localhost:8000/api/status/db`

## 📊 Health Monitoring

### System Health Endpoints

- **Basic Health Check**: `GET /`
  ```javascript
  {"status": "API is running"}
  ```

- **Detailed Health Info**: `GET /api/health`
  ```javascript
  {
    "status": "operational",
    "timestamp": "2025-11-20T10:30:00.000Z",
    "serverInfo": {
      "uptime": 3600,  // seconds
      "memoryUsage": {...},
      "nodeVersion": "v18.17.0"
    },
    "database": {
      "connected": true,
      "name": "eye-glaze-db"
    }
  }
  ```

- **Database Status**: `GET /api/status/db`
  ```javascript
  {
    "status": "success",
    "message": "Connected to MongoDB",
    "connectionState": "connected"
  }
  ```

## 🔐 Security Considerations

### Current Security Status
⚠️ **This is a development/prototype version with basic security:**
- Passwords are stored in plain text (not recommended for production)
- No JWT tokens for session management
- Basic input validation
- No rate limiting

### Production Recommendations
- Implement password hashing (bcrypt)
- Add JWT authentication
- Input validation and sanitization
- Rate limiting middleware
- HTTPS enforcement
- CORS configuration
- Environment variable validation

## 🐛 Error Handling

The API uses consistent error response format:

```javascript
{
  "status": "error",
  "message": "Descriptive error message"
}
```

### Common Error Scenarios
- **400**: Bad Request (missing required fields)
- **401**: Unauthorized (invalid credentials)
- **404**: Not Found (user/resource doesn't exist)
- **409**: Conflict (username already exists)
- **500**: Internal Server Error (database/server issues)

## 📈 Data Flow Example

Here's how a typical stress analysis session works:

1. **User Registration/Login**
   ```
   POST /api/auth/register → Create user account
   POST /api/auth/login → Authenticate user
   ```

2. **Image Upload**
   ```
   POST /api/upload/eye-image → Upload eye image to Cloudinary
   ```

3. **Analysis Submission**
   ```
   POST /api/analysis/submit → Submit analysis results
   (This includes pupilDilation, tensionRings, hasStress)
   ```

4. **Data Retrieval**
   ```
   GET /api/analysis/user/john_doe → Get all analyses
   GET /api/analysis/latest/john_doe → Get latest analysis
   ```

5. **AI Recommendations**
   ```
   GET /api/recommendations/user/john_doe → Get personalized advice
   ```

## 🧪 Testing the API

You can test the API using tools like:
- **Postman** (recommended for beginners)
- **curl** commands
- **VS Code REST Client**
- **Insomnia**

### Example Test Sequence
1. Register a new user
2. Login with those credentials
3. Upload an eye image
4. Submit an analysis result
5. Retrieve the analysis
6. Get AI recommendations

## 🤝 Integration Notes

This backend is designed to work with:
- **Frontend Applications** (React, Vue, Angular, etc.)
- **Mobile Apps** (React Native, Flutter, etc.)
- **AI/ML Models** that perform the actual eye analysis
- **External Services** (Cloudinary for images, Google Gemini for AI)

The API follows REST principles and returns JSON responses, making it compatible with any client that can make HTTP requests.

## 📝 Development Notes

- Built with **ES6 modules** (using `import`/`export`)
- Uses **async/await** for handling asynchronous operations
- Implements **error handling** with try-catch blocks
- Follows **RESTful API** conventions
- Uses **MongoDB aggregation** for data analysis
- Implements **file upload** with temporary storage and cloud upload

---

This project provides a solid foundation for a stress detection system based on eye analysis. The modular structure makes it easy to extend with additional features like more sophisticated authentication, advanced analytics, or integration with other health monitoring systems.
