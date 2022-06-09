ffmpeg -r 1 -loop 1 -i bg.jpg -i 1.mp3 -acodec copy -r 1 -shortest -vf scale=1280:720 out.flv



ffmpeg -i ./big.mp4 -vcodec libx264 -pix_fmt yuv420p -preset medium -r 30 -g 60 -b:v 2500k \
    -acodec libmp3lame -ar 44100 -threads 6 -b:a 712000 -bufsize 512k \
    -f flv rtmp://a.rtmp.youtube.com/live2/p8yp-rvmq-da4s-ky27-7m82





ffmpeg -loop 1 -y -i bg.jpg -i 1.mp3 -shortest -pix_fmt yuv420p out.mp4

ffmpeg -y -i ./bg.jpg -i ./1.mp3 -vcodec libx264 -pix_fmt yuv420p -preset medium -r 30 -g 60 -b:v 2500k \
    -acodec libmp3lame -ar 44100 -threads 6 -b:a 712000 -bufsize 512k \
    -f mp4 rtmp://a.rtmp.youtube.com/live2/p8yp-rvmq-da4s-ky27-7m82 



Works!!!!

ffmpeg -y -loop 1 -i ./bg.jpg -i ./1.mp3 -preset medium -r 30 -g 60  -vcodec h264_nvenc \
    -acodec copy -ar 44100 -threads 1  \
    -f flv rtmp://a.rtmp.youtube.com/live2/p8yp-rvmq-da4s-ky27-7m82


ffmpeg -y -loop 1 -i ./bg.jpg -i ./2.mp3 -preset medium -r 30 -g 60  -vcodec h264_nvenc -maxrate 2500k -b:a 700 -bufsize 512k -b:v 500k \
    -acodec copy -ar 44100 -threads 1  \
    -f flv rtmp://a.rtmp.youtube.com/live2/p8yp-rvmq-da4s-ky27-7m82
    
-y -i 1.mp3 -i bg.jpg -aspect 1.777778 -s 1280x720 -r 30 -ar 44100 -c:v h264_nvenc -b:v 10k -maxrate 2500k -c:a copy -b:a 700 -bufsize 512000k -preset medium -f flv -hls_list_size 0 rtmp://a.rtmp.youtube.com/live2/p8yp-rvmq-da4s-ky27-7m82

