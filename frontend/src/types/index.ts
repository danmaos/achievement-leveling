export interface User {
  id: string
  google_id: string
  email: string
  name: string
  picture: string
  xp: number
  level: number
  created_at: string
  updated_at: string
}

export interface AchievementCategory {
  id: string
  name: string
  icon: string
  color: string
}

export interface Achievement {
  id: string
  category_id: string
  category: AchievementCategory
  title: string
  description: string
  xp_reward: number
  icon: string
  difficulty: 'easy' | 'medium' | 'hard' | 'epic'
  repeatable: boolean
}

export interface UserAchievement {
  id: string
  user_id: string
  achievement_id: string
  achievement: Achievement
  completed_at: string
  times_completed: number
}

export interface CompleteResult {
  user: User
  achievement: Achievement
  xp_gained: number
  leveled_up: boolean
  new_level?: LevelThreshold
}

export interface LevelThreshold {
  level: number
  xp_required: number
  title: string
}
