# **API Documentation**
### Basic Interfaces

These are the fundamental features of Douyin (Chinese version of TikTok) implemented in the backend. It supports all users to browse Douyin videos, allows users to register accounts, publish their own videos, and lets others view these published videos.

--- 

**/douyin/feed/ - Video Feed Interface**

This interface provides a video feed without requiring a login. It returns a list of videos in reverse chronological order based on submission time. The number of videos in a single response is controlled by the server, with a maximum of 30 videos.

**Endpoint**: GET 

**Interface Definition**
```
syntax = "proto2";
package douyin.core;

message douyin_feed_request {
  optional int64 latest_time = 1;     // Optional parameter to limit the latest submission 
                                      // time of returned videos, accurate to seconds. 
                                      // Not filling it means the current time.
  optional string token = 2;          // Optional parameter for logged-in users.
}

message douyin_feed_response {
  required int32 status_code = 1;     // Status code. 0 for success, other values for failure.
  optional string status_msg = 2;     // Status description.
  repeated Video video_list = 3;      // List of videos.
  optional int64 next_time = 4;       // The earliest submission time among the videos returned this time, 
                                      // used as 'latest_time' for the next request. 
}

message Video {
  required int64 id = 1;                // Unique identifier for the video.
  required User author = 2;             // Video author information.
  required string play_url = 3;         // Video playback URL.
  required string cover_url = 4;        // Video cover image URL.
  required int64 favorite_count = 5;    // Total number of likes for the video
  required int64 comment_count = 6;     // Total number of comments on the video
  required bool is_favorite = 7;        // true - Liked, false - Not liked
  required string title = 8;            // Video title
}

message User {
  required int64 id = 1;                    // User ID
  required string name = 2;                 // User name
  optional int64 follow_count = 3;          // Total number of followings
  optional int64 follower_count = 4;        // Total number of followers
  required bool is_follow = 5;              // true - Follwing，false - Not following
  optional string avatar = 6;               // User avatar
  optional string background_image = 7;     // User profile page banner image
  optional string signature = 8;            // Personal bio
  optional int64 total_favorited = 9;       // Total number of received likes
  optional int64 work_count = 10;           // Number of works (videos)
  optional int64 favorite_count = 11;       // Number of likes given
}
```

--- 

**/douyin/user/register/ - User Registration Interface**

This interface is used for new user registration. Users provide a username, password, and nickname. The username must be unique. After successful registration, the user ID and authentication token are returned.

**Endpoint:** POST

**Interface Definition**
```
syntax = "proto2";
package douyin.core;

message douyin_user_register_request {
  required string username = 1;   // Registration username, maximum length of 32 characters.
  required string password = 2;   // Password, maximum length of 32 characters.
}

message douyin_user_register_response {
  required int32 status_code = 1;   // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;   // Status description.
  required int64 user_id = 3;       // User ID.
  required string token = 4;        // User authentication token.
}
```

--- 

**/douyin/user/login/ - User Login Interface**

This interface is used for user login using a username and password. After successful login, the user ID and authentication token are returned.

**Endpoint:** POST

**Interface Definition**
```
syntax = "proto2";
package douyin.core;

message douyin_user_login_request {
  required string username = 1;   // Login username.
  required string password = 2;   // Login password.
}

message douyin_user_login_response {
  required int32 status_code = 1;     // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;     // Status description.
  required int64 user_id = 3;         // User ID.
  required string token = 4;          // User authentication token.
}
```

---

**/douyin/user/ - User Information Interface**

This interface is used to retrieve information about the logged-in user, including their ID, nickname. If the interaction functionality is implemented, also return the number of followings and followers.

**Endpoint:** GET

**Interface Definition**
```
syntax = "proto2";
package douyin.core;

message douyin_user_request {
  required int64 user_id = 1;   // User ID.
  required string token = 2;    // User authentication token.
}

message douyin_user_response {
  required int32 status_code = 1;   // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;   // Status description.
  required User user = 3;           // User information.
}

message User {
  required int64 id = 1;                  // User ID.
  required string name = 2;               // User name.
  optional int64 follow_count = 3;        // Total number of followings.
  optional int64 follower_count = 4;      // Total number of followers.
  required bool is_follow = 5;            // true - Following, false - Not following.
  optional string avatar = 6;             // User avatar.
  optional string background_image = 7;   // User profile page banner image.
  optional string signature = 8;          // Personal bio.
  optional int64 total_favorited = 9;     // Total number of received likes.
  optional int64 work_count = 10;         // Number of works (videos).
  optional int64 favorite_count = 11;     // Number of likes given.
}
```

--- 

**/douyin/publish/action/ -  Video Publishing Interface**

This interface allows a logged-in user to upload a video.

**Endpoint:** POST

**Interface Definition**
```
syntax = "proto2";
package douyin.core;

message douyin_publish_action_request {
  required string token = 1;    // User authentication token.
  required bytes data = 2;      // Video data.
  required string title = 3;    // Video title.
}

message douyin_publish_action_response {
  required int32 status_code = 1;   // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;   // Status description.
}
```

---

**/douyin/publish/list/ - Published Video List Interface**

This interface provides a list of videos published by the logged-in user, displaying all submitted videos.

**Endpoint:** GET

**Interface Definition**
```
syntax = "proto2";
package douyin.core;

message douyin_publish_list_request {
  required int64 user_id = 1;     // User ID.
  required string token = 2;      // User authentication token.
}

message douyin_publish_list_response {
  required int32 status_code = 1;     // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;     // Status description.
  repeated Video video_list = 3;      // List of videos published by the user.
}

message Video {
  required int64 id = 1;                // Unique identifier for the video.
  required User author = 2;             // Video author information.
  required string play_url = 3;         // Video playback URL.
  required string cover_url = 4;        // Video cover image URL.
  required int64 favorite_count = 5;    // Total number of likes for the video.
  required int64 comment_count = 6;     // Total number of comments on the video.
  required bool is_favorite = 7;        // true - Liked, false - Not liked.
  required string title = 8;            // Video title.
}

message User {
  required int64 id = 1;                    // User ID.
  required string name = 2;                 // User name.
  optional int64 follow_count = 3;          // Total number of followings.
  optional int64 follower_count = 4;        // Total number of followers.
  required bool is_follow = 5;              // true - Following, false - Not following.
  optional string avatar = 6;               // User avatar.
  optional string background_image = 7;     // User profile page banner image.
  optional string signature = 8;            // Personal bio.
  optional int64 total_favorited = 9;       // Total number of received likes.
  optional int64 work_count = 10;           // Number of works (videos).
  optional int64 favorite_count = 11;       // Number of likes given.
}
```

---

### Interaction Interfaces

These interfaces enable user interactions, such as liking videos and commenting on them. Users can view comment lists, but only logged-in users can post comments.

---

**/douyin/favorite/action/ - Like Action Interface**

This interface allows a logged-in user to like or unlike a video.

**Endpoint:** POST

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.first;

message douyin_favorite_action_request {
  required string token = 1;          // User authentication token.
  required int64 video_id = 2;        // Video ID.
  required int32 action_type = 3;     // 1 - Like, 2 - Unlike.
}

message douyin_favorite_action_response {
  required int32 status_code = 1;     // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;     // Status description.
}
```

---

**/douyin/favorite/list/ - Liked Video List Interface**

This interface provides a list of videos that a logged-in user has liked.

**Endpoint:** GET

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.first;

message douyin_favorite_list_request {
  required int64 user_id = 1;     // User ID.
  required string token = 2;      // User authentication token.
}

message douyin_favorite_list_response {
  required int32 status_code = 1;     // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;     // Status description.
  repeated Video video_list = 3;      // List of videos liked by the user.
}

message Video {
  required int64 id = 1;                // Unique identifier for the video.
  required User author = 2;             // Video author information.
  required string play_url = 3;         // Video playback URL.
  required string cover_url = 4;        // Video cover image URL.
  required int64 favorite_count = 5;    // Total number of likes for the video.
  required int64 comment_count = 6;     // Total number of comments on the video.
  required bool is_favorite = 7;        // true - Liked, false - Not liked.
  required string title = 8;            // Video title.
}

message User {
  required int64 id = 1;                    // User ID.
  required string name = 2;                 // User name.
  optional int64 follow_count = 3;          // Total number of followings.
  optional int64 follower_count = 4;        // Total number of followers.
  required bool is_follow = 5;              // true - Following, false - Not following.
  optional string avatar = 6;               // User avatar.
  optional string background_image = 7;     // User profile page banner image.
  optional string signature = 8;            // Personal bio.
  optional int64 total_favorited = 9;       // Total number of received likes.
  optional int64 work_count = 10;           // Number of works (videos).
  optional int64 favorite_count = 11;       // Number of likes given.
}
```

---

**/douyin/comment/action/ - Comment Action Interface**

This interface allows a logged-in user to post or delete comments on a video.

**Endpoint:** POST

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.first;

message douyin_comment_action_request {
  required string token = 1;            // User authentication token.
  required int64 video_id = 2;          // Video ID.
  required int32 action_type = 3;       // 1 - Post comment, 2 - Delete comment.
  optional string comment_text = 4;     // Comment text provided by the user, used when action_type=1.
  optional int64 comment_id = 5;        // Comment ID to delete, used when action_type=2.
}

message douyin_comment_action_response {
  required int32 status_code = 1;     // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;     // Status description.
  optional Comment comment = 3;       // Comment content returned for successful comment posting, 
                                      // no need to fetch the entire list again.
}

message Comment {
  required int64 id = 1;              // Comment ID.
  required User user = 2;             // Commenting user information.
  required string content = 3;        // Comment content.
  required string create_date = 4;    // Comment creation date in the format mm-dd.
}
```

---

**/douyin/comment/list/ - Video Comment List Interface**

This interface allows users to view all comments on a video, ordered by reverse chronological order based on posting time.

**Endpoint:** GET

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.first;

message douyin_comment_list_request {
  required string token = 1;      // User authentication token.
  required int64 video_id = 2;    // Video ID.
}

message douyin_comment_list_response {
  required int32 status_code = 1;       // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;       // Status description.
  repeated Comment comment_list = 3;    // Comment list.
}

message Comment {
  required int64 id = 1;              // Comment ID.
  required User user = 2;             // Commenting user information.
  required string content = 3;        // Comment content.
  required string create_date = 4;    // Comment creation date in the format mm-dd.
}

message User {
  required int64 id = 1;                    // User ID.
  required string name = 2;                 // User name.
  optional int64 follow_count = 3;          // Total number of followings.
  optional int64 follower_count = 4;        // Total number of followers.
  required bool is_follow = 5;              // true - Following, false - Not following.
  optional string avatar = 6;               // User avatar.
  optional string background_image = 7;     // User profile page banner image.
  optional string signature = 8;            // Personal bio.
  optional int64 total_favorited = 9;       // Total number of received likes.
  optional int64 work_count = 10;           // Number of works (videos).
  optional int64 favorite_count = 11;       // Number of likes given.
}

```

--- 

### Relation Interfaces

These interfaces handle user relationships, allowing users to follow and unfollow each other. Logged-in users can see their own following and follower lists.

--- 

**/douyin/relation/action/ - Relationship Action Interface**

This interface allows a logged-in user to follow or unfollow another user.

**Endpoint:** POST

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.second;

message douyin_relation_action_request {
  required string token = 1;          // User authentication token.
  required int64 to_user_id = 2;      // Target user ID.
  required int32 action_type = 3;     // 1 - Follow, 2 - Unfollow.
}

message douyin_relation_action_response {
  required int32 status_code = 1; // Status code, 0 for success, other values for failure.
  optional string status_msg = 2; // Status description.
}
```

---

**/douyin/relatioin/follow/list/ - User Following List Interface**

This interface provides a list of users that a logged-in user is following.

**Endpoint:** GET

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.second;

message douyin_relation_follow_list_request {
  required int64 user_id = 1;     // User ID.
  required string token = 2;      // User authentication token.
}

message douyin_relation_follow_list_response {
  required int32 status_code = 1;     // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;     // Status description.
  repeated User user_list = 3;        // List of user information.
}

message User {
  required int64 id = 1;                    // User ID.
  required string name = 2;                 // User name.
  optional int64 follow_count = 3;          // Total number of followings.
  optional int64 follower_count = 4;        // Total number of followers.
  required bool is_follow = 5;              // true - Following, false - Not following.
  optional string avatar = 6;               // User avatar.
  optional string background_image = 7;     // User profile page banner image.
  optional string signature = 8;            // Personal bio.
  optional int64 total_favorited = 9;       // Total number of received likes.
  optional int64 work_count = 10;           // Number of works (videos).
  optional int64 favorite_count = 11;       // Number of likes given.
}
```

---

**/douyin/relation/follower/list/ - User Follower List Interface**

This interface provides a list of users who follow a logged-in user.

**Endpoint:** GET

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.second;

message douyin_relation_follower_list_request {
  required int64 user_id = 1;     // User ID.
  required string token = 2;      // User authentication token.
}

message douyin_relation_follower_list_response {
  required int32 status_code = 1;     // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;     // Status description.
  repeated User user_list = 3;        // List of user information.
}

message User {
  required int64 id = 1;                    // User ID.
  required string name = 2;                 // User name.
  optional int64 follow_count = 3;          // Total number of followings.
  optional int64 follower_count = 4;        // Total number of followers.
  required bool is_follow = 5;              // true - Following, false - Not following.
  optional string avatar = 6;               // User avatar.
  optional string background_image = 7;     // User profile page banner image.
  optional string signature = 8;            // Personal bio.
  optional int64 total_favorited = 9;       // Total number of received likes.
  optional int64 work_count = 10;           // Number of works (videos).
  optional int64 favorite_count = 11;       // Number of likes given.
}
```

---

**/douyin/relation/friend/list/ - User Friend List Interface**

This interface provides a list of users who are friends of a logged-in user.

**Endpoint:** GET

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.second;

message douyin_relation_friend_list_request {
  required int64 user_id = 1;     // User ID.
  required string token = 2;      // User authentication token.
}

message douyin_relation_friend_list_response {
  required int32 status_code = 1;       // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;       // Status description.
  repeated FriendUser user_list = 3;    // List of user information.
}

message User {
  required int64 id = 1;                    // User ID.
  required string name = 2;                 // User name.
  optional int64 follow_count = 3;          // Total number of followings.
  optional int64 follower_count = 4;        // Total number of followers.
  required bool is_follow = 5;              // true - Following, false - Not following.
  optional string avatar = 6;               // User avatar.
  optional string background_image = 7;     // User profile page banner image.
  optional string signature = 8;            // Personal bio.
  optional int64 total_favorited = 9;       // Total number of received likes.
  optional int64 work_count = 10;           // Number of works (videos).
  optional int64 favorite_count = 11;       // Number of likes given.
}

message FriendUser extends User {
    optional string message = 1;    // Most recent chat message with this friend.
    required int64 msgType = 2;     // Type of message: 
                                    // 0 => Message received by the current requesting user, 
                                    // 1 => Message sent by the current requesting user.
}
```

--- 

### Messaging Interfaces

The client queries message records from the server through periodic polling.

---

**/douyin/message/chat/ - Chat Records**

Chat message records between the currently logged-in user and a specified user.

**Endpoint:** GET

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.second;

message douyin_message_chat_request {
  required string token = 1;          // User authentication token.
  required int64 to_user_id = 2;      // ID of the other user.
  required int64 pre_msg_time = 3;    // Timestamp of the latest previous message (new field - added in APK update).
}

message douyin_message_chat_response {
  required int32 status_code = 1;       // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;       // Status description.
  repeated Message message_list = 3;    // List of messages.
}

message Message {
  required int64 id = 1;              // Message ID.
  required int64 to_user_id = 2;      // ID of the message recipient.
  required int64 from_user_id = 3;    // ID of the message sender.
  required string content = 4;        // Message content.
  optional string create_time = 5;    // Message creation time.
}
```

---

**/douyin/message/action/ - Message Actions**

Actions related to messages for a logged-in user. Currently, only message sending is supported.

**Endpoint:** POST

**Interface Definition**
```
syntax = "proto2";
package douyin.extra.second;

message douyin_relation_action_request {
  required string token = 1;          // User authentication token.
  required int64 to_user_id = 2;      // ID of the other user.
  required int32 action_type = 3;     // 1 - Send message.
  required string content = 4;        // Message content.
}

message douyin_relation_action_response {
  required int32 status_code = 1;     // Status code, 0 for success, other values for failure.
  optional string status_msg = 2;     // Status description.
}
```