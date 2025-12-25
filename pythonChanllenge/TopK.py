from collections import deque

class ErrorRateMonitor:
    def __init__(self, window_size=60):
        self.window = window_size
        self.logs = deque()  # 👈 用于存储 (timestamp, is_error)
        self.err_count = 0
    
    def add(self, timestamp: int, is_error: bool):
        self.logs.append((timestamp, is_error))
        if is_error:
            self.err_count += 1
        

    def get_error_rate(self, current_time: int) -> float:
        while self.logs and current_time - self.logs[0][0] > self.window:
            _, error = self.logs.popleft()
            if error:
                self.err_count -= 1
        return self.err_count/len(self.logs)



# ========== ✅ 示例测试代码 ==========
monitor = ErrorRateMonitor()

# 添加数据
monitor.add(100, False)
monitor.add(101, True)
monitor.add(120, True)

# 查询错误率
print("Error rate at t=160:", monitor.get_error_rate(160))  # 期望输出: 1.0

monitor.add(161, False)
monitor.add(162, False)

print("Error rate at t=170:", monitor.get_error_rate(170))  # 期望
