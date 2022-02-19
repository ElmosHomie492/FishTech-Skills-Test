# FishTech-Skills-Test
## Notes For Reviewers
I had a lot of trouble trying to hook everything up to GCP. I decided to make the switch to AWS after a long struggle. After the digging that I had to do during that, the transition to AWS was a breeze.

I kept the info that is being pulled back to a minimum. My API simply pulls back GeoIP info. I had a function created to pull back WhoIs info as well. But the format that was being returned wasn't super usable for what I was trying to do. Instead of spending time trying to manually parse that data, I went ahead and just skipped it.

Terraform is incredibly powerful. I had never heard of it before speaking with Jessica during my first interview. But it seems fairly intuitive. The main struggle for me was just related to the cloud platform I started with.

No matter how this goes, I'm looking forward to exploring Terraform more.

URL for API: https://slmxgrqrwd.execute-api.us-east-1.amazonaws.com/v1/ip_info?ip={insert_ip_here}

Example: `https://slmxgrqrwd.execute-api.us-east-1.amazonaws.com/v1/ip_info?ip=199.34.228.188`
