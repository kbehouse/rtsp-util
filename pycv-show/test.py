import cv2
# cap = cv2.VideoCapture("rtsp://root:qeek@192.168.0.54:554/axis-media/media.amp")
cap = cv2.VideoCapture("rtsp://admin:123456@192.168.0.55:554/0/video1")

while(cap.isOpened()):
    ret, frame = cap.read()
    cv2.imshow('frame', frame)
    if cv2.waitKey(20) & 0xFF == ord('q'):
        break
cap.release()
cv2.destroyAllWindows()